import axios from 'axios';

export interface QueryParmas {
  page: number;
  limit: number;
  [key: string]: any;
}
export interface HostsExtendType {
	id?:number | undefined;
	host?:string;
	cloudType:string;
	startTime:string;
	endTime:string;
	izCrond:number;
	remarks:string;
}

export function queryHostExtendList(params: QueryParmas) {
  return axios.post('/api/hostExtent/list', params);
}
export function deleteHostExtend(id: any) {
  return axios.delete(`/api/hostExtent/delete?id=${id}`);
}

export function hostExtendInsert(data: HostsExtendType) {
  return axios.post('/api/hostExtent/insert', data);
}
export function hostExtendUpdate(data: HostsExtendType) {
  return axios.post('/api/hostExtent/update', data);
}
export function homeCount() {
  return axios.get('/api/home/count', {});
}
