interface ResponseData<T = never> {
  code: number;
  data?: T;
  msg: string;
}

interface PaginationData<T> {
  keyword?: string;
  current: number;
  total: number;
  size: number;
  list: T;
}

interface AmapDistrictData {
  name: string;
  adcode: string;
  center: string;
  citycode: string;
  districts: AmapDistrictData[];
  level: string;
}

interface AmapParams {
  key: string;
  subdistrict: number;
}

interface AmapJSON {
  status: string;
  info: string;
  infocode: string;
  suggestion: {
    keywords: string[];
    cities: string[];
  };
  count?: string;
  districts: AmapDistrictData[];
}

interface GeoFeatureData {
  type: string;
  properties: {
    acroutes: number[];
    adcode: number | string;
    center: number[];
    centroid: number[];
    childrenNum: number;
    level: string;
    name: string;
    parent: {
      adcode: number;
    };
    subFeatureIndex: number;
  };
  geometry: {
    type: string;
    coordinates: number[][][];
  };
}

interface GeoParams {
  code: string;
}

interface GeoJSON {
  type: string;
  features: GeoFeatureData[];
}

interface TimeRangeParams {
  start_time: string;
  end_time: string;
}

export type {
  ResponseData,
  PaginationData,
  AmapParams,
  AmapJSON,
  AmapDistrictData,
  GeoParams,
  GeoJSON,
  GeoFeatureData,
  TimeRangeParams,
};
