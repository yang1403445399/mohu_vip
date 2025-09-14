<script setup lang="ts">
import { useRouter } from "vue-router";
import { ref, reactive, onMounted } from "vue";
import { User, Lock, Phone, EditPen } from "@element-plus/icons-vue";
import { ElMessage } from "element-plus";
import { reqUserLoginSubmit, reqUserLoginVerify } from "@/api/user";
import type { FormInstance, FormRules } from "element-plus";
import type { ResponseData } from "@/types/common";
import type { UserLoginData, UserData } from "@/types/user";

const router = useRouter();
const loading = ref<boolean>(true);
const tabsCurrent = ref<number>(1);
const autoLogin = ref<boolean>(false);
const formRef = ref<FormInstance>();
const formData = reactive<UserLoginData>({
  username: "",
  password: "",
  mobile: "",
  code: "",
});
const formRule = reactive<FormRules<UserLoginData>>({
  username: [{ required: true, message: "请输入用户名", trigger: "blur" }],
  password: [{ required: true, message: "请输入密码", trigger: "blur" }],
  mobile: [{ required: true, message: "请输入手机号", trigger: "blur" }],
  code: [{ required: true, message: "请输入验证码", trigger: "blur" }],
});

const onUserLoginVerify = async () => {
  const response: ResponseData<UserData> = await reqUserLoginVerify();
  if (response.code === 200) {
    router.push({
      name: "home",
    });
  }
};

const onUserLoginSubmit = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;

  const valid = await formEl.validate();

  if (valid) {
    const response: ResponseData<UserData> =
      await reqUserLoginSubmit(formData);
    if (response.code === 200) {
      ElMessage({
        message: response.msg,
        type: "success",
        duration: 1500,
        onClose: () => {
          localStorage.setItem("token", response.data!.token!);
          if (autoLogin.value) {
            localStorage.setItem("login", "yes");
          }
          router.push({
            name: "home",
          });
        },
      });
    }
  }
};

const loadPageData = async () => {
  try {
    await onUserLoginVerify();
  } catch (error: any) {
    ElMessage({
      message: error.message,
      type: "error",
    });
  }
};

onMounted(async () => {
  const login = localStorage.getItem("login");

  if (login === "yes") {
    autoLogin.value = true;
    await loadPageData();
  }

  loading.value = false;
});
</script>

<template>
  <div
    class="w-full h-screen flex items-center justify-center bg-[url(@/assets/pic_001.png)] bg-center bg-cover"
    v-loading="loading"
    element-loading-text="Loading..."
  >
    <div class="w-82">
      <div class="flex justify-center items-center h-15">
        <img class="w-60" src="@/assets/logo_001.svg" alt="" />
      </div>
      <div class="mt-6">
        <el-tabs v-model="tabsCurrent" @tab-click="">
          <el-tab-pane label="账户密码登录" :name="1"></el-tab-pane>
          <el-tab-pane label="手机号登录" :name="2"></el-tab-pane>
        </el-tabs>
      </div>
      <div class="mt-2">
        <el-form ref="formRef" :model="formData" :rules="formRule">
          <el-form-item prop="username" v-if="tabsCurrent === 1">
            <el-input
              v-model="formData.username"
              placeholder="请输入用户名"
              :prefix-icon="User"
            />
          </el-form-item>
          <el-form-item prop="password" v-if="tabsCurrent === 1">
            <el-input
              v-model="formData.password"
              placeholder="请输入密码"
              :prefix-icon="Lock"
              show-password
            />
          </el-form-item>
          <el-form-item prop="mobile" v-if="tabsCurrent === 2">
            <el-input
              v-model="formData.mobile"
              placeholder="请输入手机号"
              :prefix-icon="Phone"
            />
          </el-form-item>
          <el-form-item prop="code" v-if="tabsCurrent === 2">
            <div class="w-full flex items-center justify-between">
              <el-input
                class="mr-2"
                v-model="formData.code"
                placeholder="请输入验证码"
                :prefix-icon="EditPen"
              />
              <el-button>获取验证码</el-button>
            </div>
          </el-form-item>
          <el-form-item>
            <div class="w-full flex justify-between items-center">
              <el-checkbox v-model="autoLogin" label="自动登录" />
              <el-link type="primary">忘记密码?</el-link>
            </div>
          </el-form-item>
          <el-form-item>
            <el-button
              type="primary"
              style="width: 100%"
              @click="onUserLoginSubmit(formRef)"
              >登录</el-button
            >
          </el-form-item>
        </el-form>
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
:deep(.el-tabs__nav-scroll) {
  display: flex;
  justify-content: center;
}
</style>
