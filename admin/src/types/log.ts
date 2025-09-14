import type { UserData } from "./user";

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

export type { LogListParams, LogData };
