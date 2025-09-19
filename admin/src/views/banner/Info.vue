<script setup lang="ts">
import imageSelect from "@/components/ImageSelect.vue";
import { ref, reactive, onMounted } from "vue";
import { useRouter, useRoute } from "vue-router";
import { ElMessage } from "element-plus";
import type { FormInstance, FormRules } from "element-plus";
import { trimParser } from "@/utils";
import {
  reqBannerTypeList,
  reqBannerSubmit,
  reqBannerInfo,
} from "@/api/banner";
import type { ResponseData } from "@/types/common";
import type { BannerData, BannerTypeData } from "@/types/banner";

const router = useRouter();
const route = useRoute();
const loading = ref<boolean>(true);
const typeList = ref<BannerTypeData[]>([]);
const pageName = ref<string>("新增轮播图");
const formRef = ref<FormInstance>();
const formData = reactive<BannerData>({
  type_id: 1,
  name: "",
  src: "",
});
const formRule = reactive<FormRules<BannerData>>({
  type_id: [{ required: true, message: "请选择类型", trigger: "change" }],
  name: [{ required: true, message: "请输入名称", trigger: "blur" }],
  src: [{ required: true, message: "请上传图片", trigger: "blur" }],
});

const goBack = () => {
  router.back();
};

const getBannerTypeList = async () => {
  const response: ResponseData<BannerTypeData[]> = await reqBannerTypeList();
  if (response.code === 200) {
    typeList.value = response.data!;
  }
};

const getBannerInfo = async () => {
  const response: ResponseData<BannerData> = await reqBannerInfo({
    id: Number(route.query.id),
  });
  if (response.code === 200) {
    Object.assign(formData, response.data!);
  }
};

const onBannerSubmit = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;

  const valid = await formEl.validate();

  if (valid) {
    try {
      const response: ResponseData = await reqBannerSubmit(formData);
      if (response.code === 200) {
        ElMessage({
          message: "提交成功",
          type: "success",
          onClose: () => {
            goBack();
          },
        });
      }
    } catch (error: any) {
      ElMessage({
        message: error.message,
        type: "error",
      });
    }
  }
};

const onInit = async () => {
  try {
    await getBannerTypeList();
    if (route.query.id) {
      pageName.value = "编辑轮播图";
      await getBannerInfo();
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
        class="grid grid-cols-3 gap-8 items-start"
        ref="formRef"
        label-position="top"
        :model="formData"
        :rules="formRule"
      >
        <div class="grid grid-cols-1 gap-4">
          <el-form-item label="名称" prop="name">
            <el-input
              v-model="formData.name"
              :formatter="trimParser"
              placeholder="请输入名称"
            />
          </el-form-item>
          <el-form-item label="图片" prop="src">
            <component
              :is="imageSelect"
              v-model="formData.src"
              placeholder="请选择图片"
            />
          </el-form-item>
        </div>
        <div class="grid grid-cols-1 gap-4">
          <el-form-item label="链接" prop="url">
            <el-input
              v-model="formData.url"
              :formatter="trimParser"
              placeholder="请输入链接"
            />
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
        </div>
        <div class="grid grid-cols-1 gap-4">
          <el-form-item label="类型" prop="type_id">
            <el-select v-model="formData.type_id" placeholder="请选择类型">
              <el-option
                v-for="value in typeList"
                :key="value.id"
                :label="value.name"
                :value="value.id"
              />
            </el-select>
          </el-form-item>
        </div>
      </el-form>
    </div>
    <div class="flex justify-end p-4 border-t border-[var(--el-border-color)]">
      <el-button>重置</el-button>
      <el-button type="primary" @click="onBannerSubmit(formRef)"
        >保存</el-button
      >
    </div>
  </div>
</template>

<style lang="scss" scoped>
:deep(.el-form-item) {
  margin: 0;

  .el-form-item__label {
    line-height: 1;
    margin-bottom: 12px;
  }
}
</style>
