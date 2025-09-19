interface ColumnData {
  id?: number;
  parent_id: number;
  name: string;
  alias?: string;
  thumb?: string;
  keywords?: string;
  intro?: string;
  content?: string;
  sort: number;
  size: number;
  update_at?: string;
  create_at?: string;
  state?: number;
}

export type { ColumnData };