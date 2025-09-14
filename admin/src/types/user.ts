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

interface UserLoginData {
  username?: string;
  password?: string;
  mobile?: string;
  code?: string;
}

export type { UserData, UserLoginData };
