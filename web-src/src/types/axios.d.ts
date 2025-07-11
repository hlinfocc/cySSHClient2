import { AxiosRequestConfig } from 'axios';

export type ErrorMessageMode = 'none' | 'alert' | 'message'|'notify' | undefined;

export interface RequestOptions {
  apiUrl?: string;
  isJoinPrefix?: boolean;
  urlPrefix?: string;
  joinParamsToUrl?: boolean;
  formatDate?: boolean;
  isTransformResponse?: boolean;
  isReturnNativeResponse?: boolean;
  ignoreRepeatRequest?: boolean;
  joinTime?: boolean;
  withToken?: boolean;
  errorMessageMode?:ErrorMessageMode,
  timeout?: number,
  retry?: {
    count: number;
    delay: number;
  };
}

export interface Result<T = any> {
  code: number;
  data: T;
}

export interface AxiosRequestConfigRetry extends AxiosRequestConfig {
  retryCount?: number;
}
