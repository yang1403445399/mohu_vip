import router from "@/router";
import axios from "axios";
import { ElMessage } from "element-plus";
import type {
  AxiosInstance,
  AxiosResponse,
  AxiosRequestConfig,
  InternalAxiosRequestConfig,
} from "axios";


// 创建axios实例
const service: AxiosInstance = axios.create({
  baseURL: "/admin",
  timeout: 5000,
  headers: {
    "Content-Type": "application/json;charset=utf-8",
  },
});

// 请求拦截器
service.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    const token = localStorage.getItem("token");
    if (token) {
      config.headers = config.headers || {};
      config.headers["Authorization"] = token;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// 响应拦截器
service.interceptors.response.use(
  (response: AxiosResponse) => {
    const { data } = response;

    if (data.code === 400) {
      ElMessage.error(data.msg);
    }

    if (data.code === 403) {
      localStorage.clear();
      router.push({ name: "login" });
    }

    return Promise.resolve(data);
  },
  (error) => {
    return Promise.reject(error);
  }
);

const get = <T = any>(
  url: string,
  params?: object,
  config?: AxiosRequestConfig
): Promise<T> => {
  return service.get(url, { params, ...config });
};

const post = <T = any>(
  url: string,
  data?: object,
  config?: AxiosRequestConfig
): Promise<T> => {
  return service.post(url, data, config);
};

export { get, post };
