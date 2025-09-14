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

export type { ImageListParams, ImageData, ImageUploadData, ImageColumnData };
