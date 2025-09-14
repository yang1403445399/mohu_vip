import { get } from "@/utils/axios";
import type { ResponseData, PaginationData } from "@/types/common";
import type { LogListParams, LogData } from "@/types/log";

const reqLogList = (
  params: LogListParams
): Promise<ResponseData<PaginationData<LogData[]>>> => {
  return get("/log/list", params);
};

export { reqLogList };
