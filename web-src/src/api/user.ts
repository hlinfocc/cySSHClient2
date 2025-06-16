import axios from 'axios';
import { UserInfo } from '@/store/modules/user/types';
import { request } from '@/utils/request';

export interface LoginData {
  account: string;
  password: string;
}

export interface LoginDataMail {
  mailname: string;
  mailpassword: string;
}

export interface LoginRes {
  token: string;
  userInfo: UserInfo;
}
export interface UserRes {
  chartData: [];
  tableData: [];
}
export interface UserData {
  sort?: number | undefined;
  startTime?: string;
  endTime?: string;
  filterStatus?: [];
  filterType?: [];
}

export function login(data: LoginData) {
  return request.post<any>({
    url:'/api/user/login',
    data
  },{isTransformResponse:false});
}
export function loginMail(data: LoginDataMail) {
  return axios.post<LoginRes>('/api/mail/login', data);
}

export function logout() {
  return axios.post<LoginRes>('/api/user/logout');
}

export function getUserInfo() {
  return axios.get<LoginRes>(`/api/user/userInfo`);
}

export function updateUserInfo(data: UserInfo) {
  return request.post<any>({url:`/api/userInfo/update`, data},{isTransformResponse:false});
}

export function getUserData(data?: UserData) {
  return axios.post<UserRes>('/api/user/data', data);
}

export function registerUser(data: LoginData) {
  return axios.post<UserInfo>('/api/user/register', data);
}
