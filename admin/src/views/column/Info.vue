<script setup lang="ts">
import imageSelect from "@/components/ImageSelect.vue";
import tiptapEditor from "@/components/TiptapEditor.vue";
import { ref, reactive, onMounted } from "vue";
import { useRouter, useRoute } from "vue-router";
import { ElMessage } from "element-plus";
import type { FormInstance, FormRules } from "element-plus";
import { trimParser } from "@/utils";
import type { ColumnData } from "@/types/column";

const router = useRouter();
const route = useRoute();
const loading = ref<boolean>(true);
const pageName = ref<string>("新增栏目");
const formRef = ref<FormInstance>();
const formData = reactive<ColumnData>({
  parent_id: 0,
  name: "",
  sort: 0,
  size: 10,
});
const formRule = reactive<FormRules<ColumnData>>({
  parent_id: [{ required: true, message: "请选择所属栏目", trigger: "change" }],
  name: [{ required: true, message: "请输入名称", trigger: "blur" }],
  sort: [{ required: true, message: "请输入排序", trigger: "blur" }],
  size: [{ required: true, message: "请输入每页显示", trigger: "blur" }],
});

const goBack = () => {
  router.back();
};

const getColumnList = async () => {};

const getColumnInfo = async () => {};

const onColumnSubmit = async (formEl: FormInstance | undefined) => {};

const onInit = async () => {
  try {
    await getColumnList();
    if (route.query.id) {
      pageName.value = "编辑栏目";
      await getColumnInfo();
    }
  } catch (error: any) {
    ElMessage({
      message: error.message,
      type: "error",
    });
  } finally {
    loading.value = false;
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
    <div class="flex items-center justify-between p-4">
      <div class="text-base text-[var(--el-text-color-primary)] leading-none">
        {{ pageName }}
      </div>
      <div class="flex">
        <el-link class="leading-none" @click="goBack">返回列表</el-link>
      </div>
    </div>
    <div class="p-4 border-t border-[var(--el-border-color)]">
      <el-form
        class="grid grid-cols-5 gap-8 items-start"
        ref="formRef"
        label-position="top"
        :model="formData"
        :rules="formRule"
      >
        <div class="col-span-2 grid grid-cols-1 gap-4">
          <el-form-item label="上级" prop="parent_id">
            <el-select
              v-model="formData.parent_id"
              placeholder="请选择上级"
            >
              <el-option label="无上级" :value="0"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="名称" prop="name">
            <el-input
              v-model="formData.name"
              :formatter="trimParser"
              placeholder="请输入名称"
            ></el-input>
          </el-form-item>
          <el-form-item label="别名" prop="alias">
            <el-input
              v-model="formData.alias"
              :formatter="trimParser"
              placeholder="请输入别名"
            ></el-input>
          </el-form-item>
          <el-form-item label="缩略图" prop="src">
            <component
              :is="imageSelect"
              v-model="formData.thumb"
              placeholder="请选择缩略图"
            />
          </el-form-item>
          <el-form-item label="关键词" prop="keywords">
            <el-input
              v-model="formData.keywords"
              :formatter="trimParser"
              placeholder="请输入关键词"
            ></el-input>
          </el-form-item>
          <el-form-item label="描述" prop="intro">
            <el-input
              v-model="formData.intro"
              :formatter="trimParser"
              placeholder="请输入描述"
              type="textarea"
              show-word-limit
              maxlength="120"
              :rows="3"
            ></el-input>
          </el-form-item>
          <el-form-item label="排序" prop="sort">
            <el-input-number
              class="w-full!"
              v-model="formData.sort"
              :min="0"
              controls-position="right"
              placeholder="请输入排序"
            />
          </el-form-item>
          <el-form-item label="每页数量" prop="size">
            <el-input-number
              class="w-full!"
              v-model="formData.size"
              :min="0"
              controls-position="right"
              placeholder="请输入每页数量"
            />
          </el-form-item>
        </div>
        <div class="col-span-3 grid grid-cols-1 gap-4">
          <el-form-item label="内容" prop="content">
            <component :is="tiptapEditor" v-model="formData.content" />
          </el-form-item>
        </div>
      </el-form>
    </div>
    <div class="flex justify-end p-4 border-t border-[var(--el-border-color)]">
      <el-button>重置</el-button>
      <el-button type="primary" @click="onColumnSubmit(formRef)"
        >保存</el-button
      >
    </div>
  </div>
</template>

<style scoped lang="scss">
:deep(.el-form-item) {
  margin: 0;

  .el-form-item__label {
    line-height: 1;
    margin-bottom: 12px;
  }
}
</style>
