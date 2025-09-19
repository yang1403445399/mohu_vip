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
  type?: BannerTypeData;
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

export type { BannerListParams, BannerInfoParams, BannerData, BannerTypeData };
