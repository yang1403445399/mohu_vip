<script setup lang="ts">
import { onMounted, reactive } from "vue";
import { useRouter } from "vue-router";
import { ElMessage, ElMessageBox } from "element-plus";
import { reqBannerList } from "@/api/banner";
import printJS from "print-js";
import * as XLSX from "xlsx";
import { saveAs } from "file-saver";
import { Picture, Search } from "@element-plus/icons-vue";
import type { ResponseData, PaginationData } from "@/types/common";
import type { BannerData } from "@/types/banner";

const router = useRouter();

const bannerListData = reactive({
  load: true,
  data: {
    keyword: "",
    current: 1,
    total: 0,
    size: 10,
    list: [],
  } as PaginationData<BannerData[]>,
});

const getBannerList = async () => {
  const response: ResponseData<PaginationData<BannerData[]>> =
    await reqBannerList({
      current: bannerListData.data.current,
      size: bannerListData.data.size,
    });
  if (response.code === 200) {
    bannerListData.data = response.data!;
  }
  bannerListData.load = false;
};

const onBannerCurrentChange = async () => {
  bannerListData.load = true;
  await getBannerList();
};

const onBannerSizeChange = async () => {
  bannerListData.load = true;
  bannerListData.data.current = 1;
  await getBannerList();
};

const onBannerTablePrint = () => {
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

const onBannerTableExport = () => {
  // 准备导出数据
  const exportData = bannerListData.data.list.map((item) => ({
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

const onBannerAdd = () => {
  router.push({
    name: "bannerInfo",
  });
};

const onBannerEdit = (row: BannerData) => {
  router.push({
    name: "bannerInfo",
    query: {
      id: row.id,
    },
  });
};

const onBannerDelete = (row: BannerData) => {
  ElMessageBox.confirm("是否确认删除？删除后将无法恢复！", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(async () => {
    console.log("删除轮播图", row);
  });
};

const loadPageData = async () => {
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
  await loadPageData();
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
          <div class="ml-4 flex cursor-pointer" @click="onBannerTablePrint()">
            <el-icon><Printer /></el-icon>
          </div>
        </el-tooltip>
        <el-tooltip content="下载">
          <div class="ml-4 flex cursor-pointer" @click="onBannerTableExport()">
            <el-icon><Download /></el-icon>
          </div>
        </el-tooltip>
      </div>
    </div>
    <div
      class="flex items-center justify-between px-4 py-4 border-t border-[var(--el-border-color)]"
    >
      <div class="flex w-70">
        <el-input
          :prefix-icon="Search"
          v-model="bannerListData.data.keyword"
          placeholder="请输入关键词搜索"
        />
      </div>
      <div class="flex">
        <el-button type="primary" plain @click="onBannerAdd()">新增</el-button>
        <el-button type="danger" plain>删除</el-button>
      </div>
    </div>
    <div id="print-table">
      <el-table
        class="w-full"
        row-key="id"
        v-loading="bannerListData.load"
        :data="bannerListData.data.list"
      >
        <el-table-column type="selection" width="50" />
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="name" label="名称" show-overflow-tooltip />
        <el-table-column prop="src" label="图片">
          <template #default="scope">
            <el-popover
              :width="344"
              trigger="hover"
              placement="right"
              v-if="scope.row.src && scope.row.src.length > 0"
            >
              <template #reference>
                <el-button :icon="Picture" type="primary" circle />
              </template>
              <el-image
                class="w-80 h-40 block!"
                :src="scope.row.src"
                fit="cover"
              />
            </el-popover>
          </template>
        </el-table-column>
        <el-table-column prop="url" label="链接" />
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
            />
          </template>
        </el-table-column>
        <el-table-column label="操作" width="120">
          <template #default="scope">
            <el-link
              type="primary"
              class="mr-3"
              @click="onBannerEdit(scope.row)"
              >修改</el-link
            >
            <el-link type="danger" @click="onBannerDelete(scope.row)"
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
        v-model:current-page="bannerListData.data.current"
        v-model:page-size="bannerListData.data.size"
        :page-sizes="[10, 20, 40, 80]"
        :total="bannerListData.data.total"
        @current-change="onBannerCurrentChange"
        @size-change="onBannerSizeChange"
      />
    </div>
  </div>
</template>

<style lang="scss" scoped></style>
