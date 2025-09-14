import axios from "axios";
import { ElMessage } from "element-plus";
import type { AmapParams, AmapJSON, GeoParams, GeoJSON } from "@/types/common";

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

export { reqAmapJson, reqAreaJson };
