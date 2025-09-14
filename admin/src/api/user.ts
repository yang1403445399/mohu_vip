import { get, post } from "@/utils/axios";
import type { ResponseData } from "@/types/common";
import type { UserLoginData, UserData } from "@/types/user";

const reqUserLoginSubmit = (
  data: UserLoginData
): Promise<ResponseData<UserData>> => {
  return post("/user/login/submit", data);
};

const reqUserLoginVerify = (): Promise<ResponseData<UserData>> => {
  return get("/user/login/verify");
};

export { reqUserLoginSubmit, reqUserLoginVerify };
