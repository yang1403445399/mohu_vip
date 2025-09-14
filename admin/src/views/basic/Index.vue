<script setup lang="ts">
import inputImageSelect from "@/components/InputImageSelect.vue";
import { ref, reactive, onMounted } from "vue";
import { ElMessage } from "element-plus";
import type { FormInstance, FormRules } from "element-plus";
import type { ResponseData, BasicFormatData, BasicData } from "@/types";
import { reqBasicInfo, reqBasicSubmit } from "@/api";
import { entries } from "lodash";

const loading = ref<boolean>(true);
const formRef = ref<FormInstance>();
const formData = reactive<BasicFormatData>({
  site_name: "",
  site_title: "",
  site_keywords: "",
  site_intro: "",
  site_logo: "",
  site_ewm: "",
  site_mode: "1",
  site_person: "",
  site_cellphone: "",
  site_email: "",
  site_addr: "",
  site_telephone: "",
  site_icp: "",
  site_copyright: "",
  site_script: "",
});
const formRule = reactive<FormRules<BasicFormatData>>({
  site_name: [{ required: true, message: "请输入网站名称", trigger: "blur" }],
  site_title: [{ required: true, message: "请输入网站标题", trigger: "blur" }],
  site_keywords: [
    { required: true, message: "请输入网站关键词", trigger: "blur" },
  ],
  site_intro: [{ required: true, message: "请输入网站描述", trigger: "blur" }],
  site_logo: [{ required: true, message: "请输入网站logo", trigger: "blur" }],
});

const getBasicInfo = async () => {
  const response: ResponseData<BasicData[]> = await reqBasicInfo();
  if (response.code === 200) {
    response.data!.forEach((item) => {
      (formData as Record<string, any>)[item.label] = item.value;
    });
  }
};

const onBaseSubmit = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;

  const valid = await formEl.validate();

  if (valid) {
    const response: ResponseData = await reqBasicSubmit(
      entries(formData).map((item) => {
        return {
          site: "site",
          label: item[0],
          value: item[1],
        };
      })
    );
    if (response.code === 200) {
      ElMessage({
        message: response.msg,
        type: "success",
        onClose: async () => {
          loading.value = true;
          await getBasicInfo();
          setTimeout(() => {
            loading.value = false;
          }, 500);
        },
      });
    }
  }
};

const loadPageData = async () => {
  try {
    await getBasicInfo();
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
  await loadPageData();
});
</script>

<template>
  <div
    class="rounded-md border border-[var(--el-border-color)] bg-[var(--card-bg-color)]"
  >
    <div class="flex justify-between items-center p-4">
      <div class="text-base text-[var(--el-text-color-primary)] leading-none">
        基本设置
      </div>
      <div class="flex"></div>
    </div>
    <div
      class="p-4 border-t border-[var(--el-border-color)]"
      v-loading="loading"
    >
      <el-form
        class="grid grid-cols-2 gap-4"
        ref="formRef"
        label-position="top"
        :model="formData"
        :rules="formRule"
      >
        <el-form-item label="网站名称" prop="site_name">
          <el-input v-model="formData.site_name" placeholder="请输入网站名称" />
        </el-form-item>
        <el-form-item label="网站LOGO" prop="site_logo">
          <component
            :is="inputImageSelect"
            v-model="formData.site_logo"
            placeholder="请选择网站LOGO"
          />
        </el-form-item>
        <el-form-item label="网站标题" prop="site_title">
          <el-input
            v-model="formData.site_title"
            placeholder="请输入网站标题"
          />
        </el-form-item>
        <el-form-item label="网站关键词" prop="site_keywords">
          <el-input
            v-model="formData.site_keywords"
            placeholder="请输入网站关键词"
          />
        </el-form-item>
        <el-form-item label="网站描述" prop="site_intro" class="col-span-2">
          <el-input
            v-model="formData.site_intro"
            placeholder="请输入网站描述"
            show-word-limit
            type="textarea"
            maxlength="120"
            :rows="3"
          />
        </el-form-item>
        <el-form-item label="联系人" prop="site_person">
          <el-input v-model="formData.site_person" placeholder="请输入联系人" />
        </el-form-item>
        <el-form-item label="联系电话" prop="site_phone">
          <el-input
            v-model="formData.site_cellphone"
            placeholder="请输入联系电话"
          />
        </el-form-item>
        <el-form-item label="电子邮箱" prop="site_email">
          <el-input
            v-model="formData.site_email"
            placeholder="请输入电子邮箱"
          />
        </el-form-item>
        <el-form-item label="联系地址" prop="site_address">
          <el-input v-model="formData.site_addr" placeholder="请输入联系地址" />
        </el-form-item>
        <el-form-item label="座机" prop="site_telephone">
          <el-input
            v-model="formData.site_telephone"
            placeholder="请输入座机"
          />
        </el-form-item>
        <el-form-item label="模式" prop="site_mode">
          <el-radio-group v-model="formData.site_mode">
            <el-radio value="1">响应式</el-radio>
            <el-radio label="2">分离式</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="备案号" prop="site_icp">
          <el-input v-model="formData.site_icp" placeholder="请输入备案号" />
        </el-form-item>
        <el-form-item label="二维码" prop="site_ewm">
          <component
            :is="inputImageSelect"
            v-model="formData.site_ewm"
            placeholder="请选择二维码"
          />
        </el-form-item>
        <el-form-item label="所属版权" prop="site_copyright" class="col-span-2">
          <el-input
            v-model="formData.site_copyright"
            placeholder="请输入所属版权"
            type="textarea"
            show-word-limit
            maxlength="120"
            :rows="3"
          />
        </el-form-item>
        <el-form-item label="附加代码" prop="site_script" class="col-span-2">
          <el-input
            v-model="formData.site_script"
            placeholder="请输入附加代码"
            type="textarea"
            :rows="3"
          />
        </el-form-item>
      </el-form>
    </div>
    <div class="flex justify-end p-4 border-t border-[var(--el-border-color)]">
      <el-button>重置</el-button>
      <el-button type="primary" @click="onBaseSubmit(formRef)">保存</el-button>
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
