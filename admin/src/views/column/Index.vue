<script setup lang="ts">
import printJS from "print-js";
import * as XLSX from "xlsx";
import { saveAs } from "file-saver";
import { ElMessage, ElMessageBox } from "element-plus";
import { onMounted, reactive } from "vue";
import { useRouter } from "vue-router";
import { Picture } from "@element-plus/icons-vue";
import type { PaginationData } from "@/types/common";
import type { ColumnData } from "@/types/column";

const router = useRouter();
const columnTableData = reactive({
  load: true,
  data: {
    keyword: "",
    current: 1,
    total: 0,
    size: 10,
    list: [],
  } as PaginationData<ColumnData[]>,
  checked: [] as ColumnData[],
});

const getColumnList = async () => {
  columnTableData.load = false;
}

const onColumnCurrentChange = () => {};

const onColumnSizeChange = () => {};

const onColumnChecked = (row: ColumnData[]) => {
  columnTableData.checked = row;
};

const onColumnAdd = () => {
  router.push({ name: "columnInfo" });
};

const onColumnDelete = (row: ColumnData[]) => {
  ElMessageBox.confirm("是否确认删除？删除后将无法恢复！", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(() => {
    // 执行删除操作
  });
};

const onColumnSearch = () => {};

const onColumnEdit = (row: ColumnData) => {};

const onColumnStateChange = (row: ColumnData) => {};

const onColumnPrint = () => {
  printJS({
    printable: "print-table",
    type: "html",
    maxWidth: 1920,
    header: "栏目列表",
    targetStyles: ["*"],
    scanStyles: true,
    showModal: true,
  });
};

const onColumnExport = () => {
  // 准备导出数据
  const exportData = columnTableData.data.list.map((item) => ({
    ID: item.id,
    名称: item.name,
    别名: item.alias,
    缩略图: item.thumb,
    关键词: item.keywords,
    描述: item.intro,
    内容: item.content,
    排序: item.sort,
    每页显示: item.size,
    更新时间: item.update_at,
    创建时间: item.create_at,
    状态: item.state,
  }));

  // 创建工作簿和工作表
  const ws = XLSX.utils.json_to_sheet(exportData);
  const wb = XLSX.utils.book_new();
  XLSX.utils.book_append_sheet(wb, ws, "栏目列表");

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
    await getColumnList();
  } catch (error: any) {
    ElMessage({
      message: error.message,
      type: "error",
    });
  }
}

onMounted(async () => {
  await onInit();
})
</script>

<template>
  <div
    class="rounded-md border border-[var(--el-border-color)] bg-[var(--card-bg-color)]"
  >
    <div class="flex justify-between items-center p-4">
      <div class="text-base text-[var(--el-text-color-primary)] leading-none">
        栏目管理
      </div>
      <div class="flex">
        <el-tooltip content="打印">
          <div class="ml-4 flex cursor-pointer" @click="onColumnPrint()">
            <el-icon><Printer /></el-icon>
          </div>
        </el-tooltip>
        <el-tooltip content="下载">
          <div class="ml-4 flex cursor-pointer" @click="onColumnExport()">
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
            v-model="columnTableData.data.keyword"
            placeholder="请输入关键词搜索"
            @input="onColumnSearch()"
          />
        </div>
        <div class="ml-3">
          <el-button type="primary" plain @click="onColumnSearch()"
            >搜索</el-button
          >
        </div>
      </div>
      <div class="flex">
        <el-button type="primary" plain @click="onColumnAdd()">新增</el-button>
        <el-button
          type="danger"
          plain
          @click="onColumnDelete(columnTableData.checked)"
          >删除</el-button
        >
      </div>
    </div>
    <div id="print-table">
      <el-table
        class="w-full"
        row-key="id"
        v-loading="columnTableData.load"
        :data="columnTableData.data.list"
        @selection-change="onColumnChecked"
      >
        <el-table-column type="selection" width="50" />
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="name" label="名称" show-overflow-tooltip />
        <el-table-column prop="alias" label="别名" show-overflow-tooltip />
        <el-table-column label="缩略图">
          <template #default="scope">
            <el-popover
              :width="344"
              trigger="hover"
              placement="right"
              v-if="scope.row.thumb && scope.row.thumb.length > 0"
            >
              <template #reference>
                <el-button :icon="Picture" type="primary" circle />
              </template>
              <el-image
                class="w-80 h-40 block!"
                :src="scope.row.thumb"
                fit="cover"
              />
            </el-popover>
          </template>
        </el-table-column>
        <el-table-column prop="sort" label="排序" />
        <el-table-column prop="size" label="每页显示" />
        <el-table-column prop="update_at" label="更新时间" width="200" />
        <el-table-column prop="create_at" label="创建时间" width="200" />
        <el-table-column prop="state" label="状态" width="80">
          <template #default="scope">
            <el-switch
              v-model="scope.row.state"
              :active-value="1"
              :inactive-value="0"
              @change="onColumnStateChange(scope.row)"
            />
          </template>
        </el-table-column>
        <el-table-column label="操作" width="120">
          <template #default="scope">
            <el-link
              type="primary"
              class="mr-3"
              @click="onColumnEdit(scope.row)"
              >修改</el-link
            >
            <el-link type="danger" @click="onColumnDelete([scope.row])"
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
        v-model:current-page="columnTableData.data.current"
        v-model:page-size="columnTableData.data.size"
        :page-sizes="[10, 20, 40, 80]"
        :total="columnTableData.data.total"
        @current-change="onColumnCurrentChange"
        @size-change="onColumnSizeChange"
      />
    </div>
  </div>
</template>

<style scoped lang="scss"></style>
