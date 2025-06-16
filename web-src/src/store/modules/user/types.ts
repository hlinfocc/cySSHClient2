export type RoleType = '' | '*' | 'admin' | 'user';
export interface UserInfo {
  id?:number;  
  realName: string;
  account: string;
  status: number;
  userType: number;
  role?: string;
  lastLoginTime?: string;
  lastLoginIp?: string;
  createTime?: string;
  updateTime?: string;
  passwd?: string;
  thatLoginTime?: string;
  thatLoginIp?: string;
  isdelete?: number;
}
export interface UserFilterData {
  sort?: number;
  startTime?: string;
  endTime?: string;
  filterStatus?: Array<string>;
  filterType?: Array<string>;
  submit?: boolean;
  reset?: boolean;
}
export type UserState = UserInfo & UserFilterData;
