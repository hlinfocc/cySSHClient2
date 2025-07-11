import axios, { AxiosRequestConfig, AxiosResponse,InternalAxiosRequestConfig,AxiosRequestHeaders } from 'axios';
import { Modal } from '@opentiny/vue';
import locale from '@opentiny/vue-locale';
import router from '@/router';
import { getToken, clearToken } from '@/utils/auth';

export interface HttpResponse<T = unknown> {
  errMsg: string;
  code: string | number;
  data: T;
}

const { VITE_API_BASE_URL, VITE_BASE_API, VITE_MOCK_IGNORE } =
  import.meta.env || {};

if (VITE_API_BASE_URL) {
  axios.defaults.baseURL = VITE_API_BASE_URL;
}

const ignoreMockApiList = VITE_MOCK_IGNORE?.split(',') || [];
axios.interceptors.request.use(
  (config: InternalAxiosRequestConfig<any>) => {
    const isProxy = ignoreMockApiList.includes(config.url);
    if (isProxy) {
      config.url = config.url?.replace(VITE_BASE_API, '/api/v1');
    }
    console.log("interceptors.request config>>>>>>",config);
    const token = getToken();
    if (token) {
      if (!config.headers) {
        config.headers = {} as AxiosRequestHeaders;
      }
      config.headers.Authorization = `${token}`;
    }

    config.headers = { ...config.headers } as AxiosRequestHeaders;

    return config;
  },
  (error) => {
    // do something
    console.log("interceptors.request>>>>>>",error);
    return Promise.reject(error);
  }
);
// add response interceptors
axios.interceptors.response.use(
  (response: AxiosResponse<HttpResponse>) => {
    const res:any = response.data;
    if (res.code !== 200) {
      res.msg &&
        Modal.message({
          message: res.msg,
          status: 'error',
        });
      return Promise.reject(new Error(res.msg || 'Error'));
    }
    return res;
  },
  (error) => {
    console.log("interceptors.response>>>>>>",error);
    const { status, data } = error.response;
    
    if (status === 401) {
      clearToken();
      router.replace({ name: 'login' });
      Modal.message({
        message: locale.t('http.error.TokenExpire'),
        status: 'error',
      });
    } else {
      data.errMsg &&
        Modal.message({
          message: locale.t(`http.error.${data.errMsg}`),
          status: 'error',
        });
    }

    return Promise.reject(error);
  }
);
