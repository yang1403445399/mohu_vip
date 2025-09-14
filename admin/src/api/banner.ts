import { get, post } from "@/utils/axios";
import type { ResponseData, PaginationData } from "@/types/common";
import type {
  BannerListParams,
  BannerInfoParams,
  BannerData,
  BannerTypeData,
} from "@/types/banner";

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
  reqBannerList,
  reqBannerInfo,
  reqBannerSubmit,
  reqBannerDelete,
  reqBannerTypeList,
};
