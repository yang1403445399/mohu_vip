import { get } from "@/utils/axios";
import type { ResponseData, TimeRangeParams } from "@/types/common";
import type { BrowseCountData, BrowseRegionData } from "@/types/browse";

const reqBrowseCount = (): Promise<ResponseData<BrowseCountData[]>> => {
  return get("/browse/count");
};

const reqBrowseRegion = (
  params?: TimeRangeParams
): Promise<ResponseData<BrowseRegionData[]>> => {
  return get("/browse/region", params);
};

export { reqBrowseCount, reqBrowseRegion };
