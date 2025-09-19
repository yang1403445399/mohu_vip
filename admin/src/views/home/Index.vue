<script setup lang="ts">
import ChartLine from "@/components/ChartLine.vue";
import ChartMap from "@/components/ChartMap.vue";
import ChartBar from "@/components/ChartBar.vue";
import { onMounted, reactive } from "vue";
import { TopRight, BottomRight, Minus } from "@element-plus/icons-vue";
import { ElMessage } from "element-plus";
import { cloneDeep } from "lodash";
import printJS from "print-js";
import * as XLSX from "xlsx";
import { saveAs } from "file-saver";
import { reqAmapJson, reqAreaJson } from "@/api/common";
import { reqBrowseCount, reqBrowseRegion } from "@/api/browse";
import { reqArticleCount } from "@/api/article";
import { reqLogList } from "@/api/log";
import type {
  GeoFeatureData,
  GeoJSON,
  AmapJSON,
  ResponseData,
  PaginationData,
  TimeRangeParams,
} from "@/types/common";
import type { LogData } from "@/types/log";
import type { BrowseCountData, BrowseRegionData } from "@/types/browse";
import type { ArticleCountData } from "@/types/article";

const browseCountData = reactive({
  load: true,
  list: [{}, {}, {}, {}] as BrowseCountData[],
});

const browseRegionData = reactive({
  load: true,
  max: 0,
  geo: {} as GeoJSON,
  map: [] as BrowseRegionData[],
  bar: [] as (string | number)[][],
  range: [] as string[],
});

const articleCountData = reactive({
  load: true,
  bar: {} as ArticleCountData,
  range: [] as string[],
});

const logTableData = reactive({
  load: true,
  data: {
    current: 1,
    total: 0,
    size: 10,
    list: [],
  } as PaginationData<LogData[]>,
  time: {
    range: [] as string[],
    visible: false,
  },
});

const getAmapJson = async (code: string) => {
  try {
    const response: GeoJSON = await reqAmapJson({ code });

    response.features.forEach((item: GeoFeatureData) => {
      if (item.properties && item.properties.name === "海南省") {
        item.geometry.coordinates = item.geometry.coordinates.slice(0, 1);
      }
    });

    response.features = response.features.filter((item: GeoFeatureData) => {
      return item.properties.adcode !== "100000_JD";
    });

    browseRegionData.geo = response;
  } catch (error: any) {
    ElMessage({
      message: error.message,
      type: "error",
    });
  }
};

const getAreaJson = async () => {
  try {
    const response: AmapJSON = await reqAreaJson({
      key: "13f18176eb98b6757c452a64743d3c99",
      subdistrict: 1,
    });

    if (response.status === "1") {
      browseRegionData.map = response.districts[0].districts.map((item) => {
        return {
          rank: 0,
          name: item.name,
          value: 0,
        };
      });
    }
  } catch (error: any) {
    ElMessage({
      message: error.message,
      type: "error",
    });
  }
};

const getBrowseTrend = (trend: string) => {
  const isPositive = trend.includes("+");
  const isNegative = trend.includes("-");

  return {
    type: isPositive ? "positive" : isNegative ? "negative" : "neutral",
    color: isPositive ? "#67C23A" : isNegative ? "#F56C6C" : "#409EFF",
    icon: isPositive ? TopRight : isNegative ? BottomRight : Minus,
  };
};

const getBrowseCount = async () => {
  try {
    const response: ResponseData<BrowseCountData[]> = await reqBrowseCount();

    if (response.code === 200) {
      browseCountData.list = response.data!;
      browseCountData.load = false;

      setTimeout(async () => {
        await getAmapJson("100000_full");
        await getAreaJson();
        await getBrowseRegion();
      }, 1000);
    }
  } catch (error: any) {
    ElMessage({
      message: error.message,
      type: "error",
    });
  }
};

const getBrowseRegion = async (params?: TimeRangeParams) => {
  try {
    const response: ResponseData<BrowseRegionData[]> = await reqBrowseRegion(
      params
    );

    if (response.code === 200) {
      const regionMap = new Map(
        response.data!.map((item) => [item.name, item.value])
      );

      browseRegionData.map.forEach((item) => {
        if (regionMap.has(item.name)) {
          item.value = regionMap.get(item.name)!;
        }
      });

      browseRegionData.map
        .sort((a, b) => b.value - a.value)
        .forEach((item, index) => {
          item.rank = index + 1;
        });

      browseRegionData.max = Math.max(
        ...browseRegionData.map.map((item) => item.value)
      );

      browseRegionData.bar = cloneDeep(browseRegionData.map)
        .slice(0, 8)
        .map((item) => [item.rank, item.name, item.value])
        .reverse();

      browseRegionData.bar.unshift(["rank", "name", "value"]);

      browseRegionData.load = false;

      setTimeout(async () => {
        await getArticleCount();
      }, 1000);
    }
  } catch (error: any) {
    ElMessage({
      message: error.message,
      type: "error",
    });
  }
};

const onBrowseTimeChange = async (value: string[]) => {
  try {
    browseRegionData.load = true;
    browseRegionData.map.forEach((item) => {
      item.value = 0;
    });
    await getBrowseRegion({
      start_time: value[0],
      end_time: value[1],
    });
  } catch (error: any) {
    ElMessage({
      message: error.message,
      type: "error",
    });
  }
};

const getArticleCount = async (params?: TimeRangeParams) => {
  try {
    const response: ResponseData<ArticleCountData> = await reqArticleCount(
      params
    );

    if (response.code === 200) {
      articleCountData.bar = response.data!;

      articleCountData.load = false;
    }
  } catch (error: any) {
    ElMessage({
      message: error.message,
      type: "error",
    });
  }
};

const onArticleTimeChange = async (value: string[]) => {
  try {
    articleCountData.load = true;
    await getArticleCount({
      start_time: value[0],
      end_time: value[1],
    });
  } catch (error: any) {
    ElMessage({
      message: error.message,
      type: "error",
    });
  }
};

const getLogList = async (params?: TimeRangeParams) => {
  try {
    const response: ResponseData<PaginationData<LogData[]>> = await reqLogList({
      size: logTableData.data.size,
      current: logTableData.data.current,
      ...params,
    });

    if (response.code === 200) {
      logTableData.data = response.data!;
      logTableData.load = false;
    }
  } catch (error: any) {
    ElMessage({
      message: error.message,
      type: "error",
    });
  }
};

const onLogCurrentChange = async () => {
  logTableData.load = true;
  await getLogList();
};

const onLogSizeChange = async () => {
  logTableData.load = true;
  logTableData.data.current = 1;
  await getLogList();
};

const onLogPopoverSwitch = () => {
  logTableData.time.visible = !logTableData.time.visible;
};

const onLogTimeChange = async (value: string[]) => {
  try {
    logTableData.load = true;
    await getLogList({
      start_time: value[0],
      end_time: value[1],
    });
  } catch (error: any) {
    ElMessage({
      message: error.message,
      type: "error",
    });
  }
};

const onLogPrint = () => {
  printJS({
    printable: "print-table",
    type: "html",
    maxWidth: 1320,
    header: "用户操作记录",
    targetStyles: ["*"],
    scanStyles: true,
    showModal: true,
  });
};

const onLogExport = () => {
  // 准备导出数据
  const exportData = logTableData.data.list.map((item) => ({
    ID: item.id,
    用户: item.user?.nickname || "",
    路由: item.route,
    参数: item.params,
    备注: item.remark,
    创建时间: item.create_at,
  }));

  // 创建工作簿和工作表
  const ws = XLSX.utils.json_to_sheet(exportData);
  const wb = XLSX.utils.book_new();
  XLSX.utils.book_append_sheet(wb, ws, "用户操作记录");

  // 生成Excel文件并下载
  const excelBuffer = XLSX.write(wb, { bookType: "xlsx", type: "array" });
  const excelData = new Blob([excelBuffer], {
    type: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet;charset=UTF-8",
  });
  const fileName = "用户操作记录.xlsx";
  saveAs(excelData, fileName);

  ElMessage({
    message: "导出成功",
    type: "success",
  });
};

const onInit = async () => {
  await getBrowseCount();
  await getLogList();
};

onMounted(async () => {
  await onInit();
});
</script>

<template>
  <div class="grid grid-cols-4 gap-4">
    <div
      class="rounded-md border border-[var(--el-border-color)] p-4 bg-[var(--card-bg-color)]"
      v-for="(value, index) in browseCountData.list"
      :key="index"
    >
      <el-skeleton :loading="browseCountData.load">
        <template #template>
          <div v-loading="true" class="h-32"></div>
        </template>
        <template #default>
          <div class="flex justify-between items-start">
            <div
              class="text-base text-[var(--el-text-color-primary)] leading-none"
            >
              {{ value.name }}
            </div>
            <div
              class="text-xs text-[var(--el-text-color-regular)] leading-none"
            >
              {{ value.date }}
            </div>
          </div>
          <div class="mt-2">
            <el-statistic :value="value.count">
              <template #suffix>
                <div class="flex items-center pb-0.5">
                  <div
                    class="text-base px-1 rounded-full leading-none"
                    :style="{ color: getBrowseTrend(value.trend).color }"
                  >
                    {{ value.trend }}
                  </div>
                  <div class="flex">
                    <el-icon
                      :size="16"
                      :color="getBrowseTrend(value.trend).color"
                    >
                      <component :is="getBrowseTrend(value.trend).icon" />
                    </el-icon>
                  </div>
                </div>
              </template>
            </el-statistic>
          </div>
          <div class="flex mt-4">
            <component
              className="w-full h-14"
              :is="ChartLine"
              :yAxisData="value.value"
              :xAxisData="value.label"
            />
          </div>
        </template>
      </el-skeleton>
    </div>
    <div
      class="rounded-md overflow-hidden border border-[var(--el-border-color)] col-span-2 p-4 bg-[var(--card-bg-color)]"
    >
      <div class="flex justify-between items-start">
        <div class="text-base text-[var(--el-text-color-primary)] leading-none">
          访问量分布图
        </div>
        <div class="flex">
          <el-date-picker
            v-model="browseRegionData.range"
            unlink-panels
            type="daterange"
            range-separator="-"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
            value-format="YYYY-MM-DD"
            :clearable="false"
            @change="onBrowseTimeChange"
          />
        </div>
      </div>
      <el-skeleton :loading="browseRegionData.load">
        <template #template>
          <div v-loading="true" class="mt-2 h-80"></div>
        </template>
        <template #default>
          <div class="mt-2 flex">
            <component
              className="w-full h-80"
              :is="ChartMap"
              :geoData="browseRegionData.geo"
              :mapData="browseRegionData.map"
              :barData="browseRegionData.bar"
              :maxValue="browseRegionData.max"
            />
          </div>
        </template>
      </el-skeleton>
    </div>
    <div
      class="rounded-md border border-[var(--el-border-color)] col-span-2 p-4 bg-[var(--card-bg-color)]"
    >
      <div class="flex justify-between items-start">
        <div class="text-base text-[var(--el-text-color-primary)] leading-none">
          文章发布统计
        </div>
        <div class="flex">
          <el-date-picker
            v-model="articleCountData.range"
            type="monthrange"
            range-separator="-"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
            value-format="YYYY-MM-DD"
            :clearable="false"
            @change="onArticleTimeChange"
          />
        </div>
      </div>
      <el-skeleton :loading="articleCountData.load">
        <template #template>
          <div v-loading="true" class="mt-2 h-80"></div>
        </template>
        <template #default>
          <div class="mt-2 flex">
            <component
              className="w-full h-80"
              :is="ChartBar"
              :yAxisData="articleCountData.bar.value"
              :xAxisData="articleCountData.bar.label"
            />
          </div>
        </template>
      </el-skeleton>
    </div>
    <div
      class="rounded-md border border-[var(--el-border-color)] col-span-4 bg-[var(--card-bg-color)]"
    >
      <div class="flex justify-between items-center p-4">
        <div class="text-base text-[var(--el-text-color-primary)] leading-none">
          用户操作记录
        </div>
        <div class="flex">
          <el-tooltip content="查找">
            <div class="ml-4 flex cursor-pointer" @click="onLogPopoverSwitch()">
              <el-popover
                placement="left"
                trigger="click"
                :width="378"
                :visible="logTableData.time.visible"
              >
                <template #reference>
                  <el-icon><Search /></el-icon>
                </template>
                <el-date-picker
                  v-model="logTableData.time.range"
                  unlink-panels
                  type="daterange"
                  range-separator="-"
                  start-placeholder="开始日期"
                  end-placeholder="结束日期"
                  value-format="YYYY-MM-DD"
                  :clearable="false"
                  @change="onLogTimeChange"
                />
              </el-popover>
            </div>
          </el-tooltip>
          <el-tooltip content="打印">
            <div class="ml-4 flex cursor-pointer" @click="onLogPrint()">
              <el-icon><Printer /></el-icon>
            </div>
          </el-tooltip>
          <el-tooltip content="下载">
            <div class="ml-4 flex cursor-pointer" @click="onLogExport()">
              <el-icon><Download /></el-icon>
            </div>
          </el-tooltip>
        </div>
      </div>
      <div id="print-table">
        <el-table
          class="w-full"
          row-key="id"
          v-loading="logTableData.load"
          :data="logTableData.data.list"
        >
          <el-table-column prop="id" label="ID" width="80" />
          <el-table-column prop="user.nickname" label="用户" />
          <el-table-column prop="route" label="路由" />
          <el-table-column prop="params" label="参数" />
          <el-table-column prop="remark" label="备注" show-overflow-tooltip />
          <el-table-column
            prop="create_at"
            label="创建时间"
            width="200"
            sortable
          />
        </el-table>
      </div>
      <div class="p-4 flex justify-end">
        <el-pagination
          small
          background
          layout="total, sizes, prev, pager, next, jumper"
          v-model:current-page="logTableData.data.current"
          v-model:page-size="logTableData.data.size"
          :page-sizes="[10, 20, 40, 80]"
          :total="logTableData.data.total"
          @current-change="onLogCurrentChange"
          @size-change="onLogSizeChange"
        />
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped></style>
