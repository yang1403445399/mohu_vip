import axios from "axios";
import { ElMessage } from "element-plus";
import { get, post } from "@/utils/axios";
import type {
  ResponseData,
  PaginationData,
  userLoginData,
  UserData,
  AmapParams,
  AmapJSON,
  GeoParams,
  GeoJSON,
  MenuData,
  BrowseCountData,
  TimeRangeParams,
  BrowseRegionData,
  ArticleCountData,
  LogListParams,
  LogData,
  BasicData,
  ImageListParams,
  ImageData,
  ImageColumnData,
  ImageUploadData,
  BannerListParams,
  BannerInfoParams,
  BannerData,
  BannerTypeData,
} from "@/types";

const reqAmapJson = async (params: GeoParams): Promise<GeoJSON> => {
  const res = await axios.get(
    "https://geo.datav.aliyun.com/areas_v3/bound/geojson",
    { params }
  );

  return res.data;
};

const reqAreaJson = async (params: AmapParams): Promise<AmapJSON> => {
  const res = await axios.get("https://restapi.amap.com/v3/config/district", {
    params,
  });

  if (res.data.status !== "1") {
    ElMessage({ message: res.data.info, type: "error" });
  }

  return res.data;
};

const reqUserLoginSubmit = (
  data: userLoginData
): Promise<ResponseData<UserData>> => {
  return post("/user/login/submit", data);
};

const reqUserLoginVerify = (): Promise<ResponseData<UserData>> => {
  return get("/user/login/verify");
};

const reqMenuList = (): Promise<ResponseData<MenuData[]>> => {
  return get("/menu/list");
};

const reqBrowseCount = (): Promise<ResponseData<BrowseCountData[]>> => {
  return get("/browse/count");
};

const reqBrowseRegion = (
  params?: TimeRangeParams
): Promise<ResponseData<BrowseRegionData[]>> => {
  return get("/browse/region", params);
};

const reqArticleCount = (
  params?: TimeRangeParams
): Promise<ResponseData<ArticleCountData>> => {
  return get("/article/count", params);
};

const reqLogsList = (
  params: LogListParams
): Promise<ResponseData<PaginationData<LogData[]>>> => {
  return get("/log/list", params);
};

const reqBasicInfo = (): Promise<ResponseData<BasicData[]>> => {
  return get("/basic/info");
};

const reqBasicSubmit = (data: BasicData[]): Promise<ResponseData> => {
  return post("/basic/submit", data);
};

const reqImageList = (
  params: ImageListParams
): Promise<ResponseData<PaginationData<ImageData[]>>> => {
  return get("/image/list", params);
};

const reqImageUpload = async (data: ImageUploadData): Promise<ResponseData> => {
  const formData = new FormData();

  formData.append("image", data.image);

  formData.append("column_id", data.column_id.toString());

  return post("/image/upload", formData, {
    headers: {
      "Content-Type": "multipart/form-data",
    },
  });
};

const reqImageDelete = (data: number[]): Promise<ResponseData> => {
  return post("/image/delete", data);
};

const reqImageColumnList = (): Promise<ResponseData<ImageColumnData[]>> => {
  return get("/image/column/list");
};

const reqImageColumnSubmit = (data: ImageColumnData): Promise<ResponseData> => {
  return post("/image/column/submit", data);
};

const reqImageColumnDelete = (data: number[]): Promise<ResponseData> => {
  return post("/image/column/delete", data);
};

const reqBannerList = (
  params?: BannerListParams
): Promise<ResponseData<PaginationData<BannerData[]>>> => {
  return get("/banner/list", params);
};

const reqBannerInfo = (
  params: BannerInfoParams
): Promise<ResponseData<BannerData>> => {
  return get("/banner/info", params);
};

const reqBannerSubmit = (data: BannerData): Promise<ResponseData> => {
  return post("/banner/submit", data);
};

const reqBannerDelete = (data: number[]): Promise<ResponseData> => {
  return post("/banner/delete", data);
};

const reqBannerTypeList = (): Promise<ResponseData<BannerTypeData[]>> => {
  return get("/banner/type/list");
};

export {
  reqAmapJson,
  reqAreaJson,
  reqUserLoginSubmit,
  reqUserLoginVerify,
  reqMenuList,
  reqBrowseCount,
  reqBrowseRegion,
  reqArticleCount,
  reqLogsList,
  reqBasicInfo,
  reqBasicSubmit,
  reqImageList,
  reqImageUpload,
  reqImageDelete,
  reqImageColumnList,
  reqImageColumnSubmit,
  reqImageColumnDelete,
  reqBannerList,
  reqBannerInfo,
  reqBannerSubmit,
  reqBannerDelete,
  reqBannerTypeList,
};
