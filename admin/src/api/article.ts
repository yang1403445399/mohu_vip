import { get } from "@/utils/axios";
import type { ResponseData, TimeRangeParams } from "@/types/common";
import type { ArticleCountData } from "@/types/article";

const reqArticleCount = (
  params?: TimeRangeParams
): Promise<ResponseData<ArticleCountData>> => {
  return get("/article/count", params);
};

export { reqArticleCount };
