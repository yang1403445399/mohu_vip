// 使用默认类型参数处理T不一定存在的情况
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

interface UserData {
  id: number;
  role_id: number;
  username: string;
  password?: string;
  nickname: string;
  avatar: string;
  uuid: string;
  create_at: string;
  update_at: string;
  state: number;
  token?: string;
}

interface userLoginData {
  username?: string;
  password?: string;
  mobile?: string;
  code?: string;
}

interface MenuData {
  id: number;
  parent_id: number;
  name: string;
  path: string;
  icon: string;
  create_at: string;
  update_at: string;
  state: number;
}

interface MenuTreeData extends MenuData {
  children?: MenuTreeData[];
}

interface TimeRangeParams {
  start_time: string;
  end_time: string;
}

interface BrowseCountData {
  name: string;
  date: string;
  trend: string;
  count: number;
  label: string[];
  value: number[];
  group: {
    label: string;
    value: number;
  }[];
}

interface BrowseRegionData {
  rank: number;
  name: string;
  value: number;
}

interface ArticleCountData {
  label: string[];
  value: number[];
}

interface LogListParams {
  current: number;
  size: number;
  start_time?: string;
  end_time?: string;
}

interface LogData {
  id: number;
  user_id: number;
  user?: UserData;
  route: string;
  method: string;
  params: string;
  remark: string;
  create_at: string;
  update_at: string;
  state: number;
}

interface BasicData {
  site: string;
  label: string;
  value: string;
}

interface BasicFormatData {
  site_name: string;
  site_title: string;
  site_keywords: string;
  site_intro: string;
  site_logo: string;
  site_ewm: string;
  site_mode: string;
  site_person: string;
  site_cellphone: string;
  site_email: string;
  site_addr: string;
  site_telephone: string;
  site_icp: string;
  site_copyright: string;
  site_script: string;
}

interface ImageListParams {
  current: number;
  size: number;
  column_id: number;
}

interface ImageData {
  id: number;
  user_id: number;
  column_id: number;
  type_id: number;
  name: string;
  url: string;
  create_at: string;
  update_at: string;
  state: number;
}

interface ImageUploadData {
  image: File;
  column_id: number;
}

interface ImageColumnData {
  id?: number;
  user_id?: number;
  name: string;
  create_at?: string;
  update_at?: string;
  state?: number;
}

interface BannerListParams {
  keyword?: string;
  current: number;
  size: number;
}

interface BannerInfoParams {
  id: number;
}

interface BannerData {
  id?: number;
  type_id: number;
  name: string;
  src: string;
  url?: string;
  sort?: number;
  create_at?: string;
  update_at?: string;
  state?: number;
}

interface BannerTypeData {
  id: number;
  name: string;
  create_at?: string;
  update_at?: string;
  state?: number;
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
  UserData,
  userLoginData,
  MenuData,
  MenuTreeData,
  BrowseCountData,
  TimeRangeParams,
  BrowseRegionData,
  ArticleCountData,
  LogListParams,
  LogData,
  BasicData,
  BasicFormatData,
  ImageListParams,
  ImageData,
  ImageUploadData,
  ImageColumnData,
  BannerListParams,
  BannerInfoParams,
  BannerData,
  BannerTypeData,
};
