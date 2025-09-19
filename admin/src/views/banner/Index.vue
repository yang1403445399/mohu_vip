<script setup lang="ts">
import printJS from "print-js";
import * as XLSX from "xlsx";
import { saveAs } from "file-saver";
import { cloneDeep, debounce } from "lodash";
import { onMounted, reactive } from "vue";
import { useRouter } from "vue-router";
import { ElMessage, ElMessageBox } from "element-plus";
import { reqBannerList, reqBannerDelete, reqBannerSubmit } from "@/api/banner";
import type { ResponseData, PaginationData } from "@/types/common";
import type { BannerData } from "@/types/banner";

const router = useRouter();
const bannerTableData = reactive({
  load: true,
  data: {
    keyword: "",
    current: 1,
    total: 0,
    size: 10,
    list: [],
  } as PaginationData<BannerData[]>,
  checked: [] as BannerData[],
  preview: [] as string[],
});

const getBannerList = async () => {
  const response: ResponseData<PaginationData<BannerData[]>> =
    await reqBannerList({
      keyword: bannerTableData.data.keyword,
      current: bannerTableData.data.current,
      size: bannerTableData.data.size,
    });
  if (response.code === 200) {
    bannerTableData.preview = response.data!.list.map((item) => item.src);
    bannerTableData.data = response.data!;
  }
  bannerTableData.load = false;
};

const onBannerCurrentChange = async () => {
  bannerTableData.load = true;
  await getBannerList();
};

const onBannerSizeChange = async () => {
  bannerTableData.load = true;
  bannerTableData.data.current = 1;
  await getBannerList();
};

const onBannerChecked = (row?: BannerData[]) => {
  if (row) {
    bannerTableData.checked = row;
  }
};

const onBannerAdd = () => {
  router.push({
    name: "bannerInfo",
  });
};

const onBannerDelete = (row: BannerData[]) => {
  ElMessageBox.confirm("是否确认删除？删除后将无法恢复！", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(async () => {
    try {
      const response: ResponseData = await reqBannerDelete(
        row.map((item) => item.id!)
      );
      if (response.code === 200) {
        ElMessage({
          message: response.msg,
          type: "success",
          onClose: async () => {
            bannerTableData.checked = [];
            await getBannerList();
          },
        });
      }
    } catch (error: any) {
      ElMessage({
        message: error.message,
        type: "error",
      });
    }
  });
};

const onBannerSearch = debounce(async () => {
  bannerTableData.load = true;
  bannerTableData.data.current = 1;
  await getBannerList();
}, 500);

const onBannerEdit = (row: BannerData) => {
  router.push({
    name: "bannerInfo",
    query: {
      id: row.id,
    },
  });
};

const onBannerStateChange = async (row: BannerData) => {
  try {
    let postData = cloneDeep(row);
    delete postData.type;
    const response: ResponseData = await reqBannerSubmit(postData);
    if (response.code === 200) {
      ElMessage({
        message: response.msg,
        type: "success",
      });
    }
  } catch (error: any) {
    ElMessage({
      message: error.message,
      type: "error",
    });
  }
};

const onBannerPrint = () => {
  printJS({
    printable: "print-table",
    type: "html",
    maxWidth: 1920,
    header: "轮播图",
    targetStyles: ["*"],
    scanStyles: true,
    showModal: true,
  });
};

const onBannerExport = () => {
  // 准备导出数据
  const exportData = bannerTableData.data.list.map((item) => ({
    ID: item.id,
    名称: item.name,
    图片: item.src,
    链接: item.url,
    排序: item.sort,
    更新时间: item.update_at,
    创建时间: item.create_at,
  }));

  // 创建工作簿和工作表
  const ws = XLSX.utils.json_to_sheet(exportData);
  const wb = XLSX.utils.book_new();
  XLSX.utils.book_append_sheet(wb, ws, "轮播图");

  // 生成Excel文件并下载
  const excelBuffer = XLSX.write(wb, { bookType: "xlsx", type: "array" });
  const excelData = new Blob([excelBuffer], {
    type: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet;charset=UTF-8",
  });
  const fileName = "轮播图.xlsx";
  saveAs(excelData, fileName);

  ElMessage({
    message: "导出成功",
    type: "success",
  });
};

const onInit = async () => {
  try {
    await getBannerList();
  } catch (error: any) {
    ElMessage({
      message: error.message,
      type: "error",
    });
  }
};

onMounted(async () => {
  await onInit();
});
</script>

<template>
  <div
    class="rounded-md border border-[var(--el-border-color)] bg-[var(--card-bg-color)]"
  >
    <div class="flex justify-between items-center p-4">
      <div class="text-base text-[var(--el-text-color-primary)] leading-none">
        轮播图
      </div>
      <div class="flex">
        <el-tooltip content="打印">
          <div class="ml-4 flex cursor-pointer" @click="onBannerPrint()">
            <el-icon><Printer /></el-icon>
          </div>
        </el-tooltip>
        <el-tooltip content="下载">
          <div class="ml-4 flex cursor-pointer" @click="onBannerExport()">
            <el-icon><Download /></el-icon>
          </div>
        </el-tooltip>
      </div>
    </div>
    <div
      class="flex items-center justify-between px-4 py-4 border-t border-[var(--el-border-color)]"
    >
      <div class="flex">
        <div class="w-60">
          <el-input
            v-model="bannerTableData.data.keyword"
            placeholder="请输入关键词搜索"
            @input="onBannerSearch()"
          />
        </div>
        <div class="ml-3">
          <el-button type="primary" plain @click="onBannerSearch()"
            >搜索</el-button
          >
        </div>
      </div>
      <div class="flex">
        <el-button type="primary" plain @click="onBannerAdd()">新增</el-button>
        <el-button
          type="danger"
          plain
          @click="onBannerDelete(bannerTableData.checked)"
          >删除</el-button
        >
      </div>
    </div>
    <div id="print-table">
      <el-table
        class="w-full"
        row-key="id"
        v-loading="bannerTableData.load"
        :data="bannerTableData.data.list"
        @selection-change="onBannerChecked"
      >
        <el-table-column type="selection" width="50" />
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="name" label="名称" show-overflow-tooltip />
        <el-table-column prop="src" label="图片">
          <template #default="scope">
            <el-image
              class="w-8 h-8 block!"
              preview-teleported
              :src="scope.row.src"
              :preview-src-list="bannerTableData.preview"
              :zoom-rate="1.2"
              :max-scale="7"
              :min-scale="0.2"
              :z-index="100"
              :initial-index="scope.$index"
              fit="cover"
            />
          </template>
        </el-table-column>
        <el-table-column prop="url" label="链接" show-overflow-tooltip />
        <el-table-column prop="type.name" label="类型">
          <template #default="scope">
            <el-tag :type="scope.row.type_id === 1 ? `primary` : `success`">{{
              scope.row.type.name
            }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="sort" label="排序" />
        <el-table-column prop="update_at" label="更新时间" width="200" />
        <el-table-column prop="create_at" label="创建时间" width="200" />
        <el-table-column prop="state" label="状态" width="80">
          <template #default="scope">
            <el-switch
              v-model="scope.row.state"
              :active-value="1"
              :inactive-value="0"
              @change="onBannerStateChange(scope.row)"
            />
          </template>
        </el-table-column>
        <el-table-column label="操作" width="120">
          <template #default="scope">
            <el-link
              type="primary"
              class="mr-3"
              @click="onBannerEdit(scope.row)"
              >编辑</el-link
            >
            <el-link type="danger" @click="onBannerDelete([scope.row])"
              >删除</el-link
            >
          </template>
        </el-table-column>
      </el-table>
    </div>
    <div class="p-4 flex justify-end">
      <el-pagination
        small
        background
        layout="total, sizes, prev, pager, next, jumper"
        v-model:current-page="bannerTableData.data.current"
        v-model:page-size="bannerTableData.data.size"
        :page-sizes="[10, 20, 40, 80]"
        :total="bannerTableData.data.total"
        @current-change="onBannerCurrentChange"
        @size-change="onBannerSizeChange"
      />
    </div>
  </div>
</template>

<style lang="scss" scoped></style>
