import { get, post } from "@/utils/axios";
import type { ResponseData, PaginationData } from "@/types/common";
import type {
  ImageListParams,
  ImageData,
  ImageColumnData,
  ImageUploadData,
} from "@/types/image";

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

export {
  reqImageList,
  reqImageUpload,
  reqImageDelete,
  reqImageColumnList,
  reqImageColumnSubmit,
  reqImageColumnDelete,
};
