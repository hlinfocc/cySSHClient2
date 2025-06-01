import { defineStore } from 'pinia';
import {
  login as userLogin,
  loginMail as userLoginMail,
  getUserInfo,
  updateUserInfo,
  LoginData,
  LoginDataMail,
} from '@/api/user';
import { setToken, clearToken } from '@/utils/auth';
import { removeRouteListener } from '@/utils/route-listener';
import { UserState, UserInfo } from './types';

// 从 localStorage 加载初始状态
const loadInitialState = (): UserState => {
  const savedState = localStorage.getItem('USER_INFO_STORE_STATE');
  return savedState ? JSON.parse(savedState) : {} as any;
};

const useUserStore = defineStore('user', {
  state: (): UserState => loadInitialState(),

  getters: {
    userInfo(state: UserState): UserState {
      return state;
    },
  },

  actions: {
    // 保存状态到 localStorage
    saveState() {
      localStorage.setItem('USER_INFO_STORE_STATE', JSON.stringify(this.$state));
    },

    switchRoles() {
      return new Promise((resolve) => {
        this.role = this.role === 'user' ? 'admin' : 'user';
        this.saveState();
        resolve(this.role);
      });
    },

    // Set user's information
    setInfo(partial: Partial<UserState>) {
      this.$patch(partial);
      this.saveState();
    },

    // Reset user's information
    resetInfo() {
      this.$reset();
      localStorage.removeItem('USER_INFO_STORE_STATE');
    },

    // Reset filter information
    resetFilterInfo() {
      this.startTime = '';
      this.endTime = '';
      this.filterStatus = [];
      this.filterType = [];
      this.saveState();
    },

    // Get user's information
    async info() {
      const res = await getUserInfo();
      console.log("info:", res);
      if (res) {
        this.setInfo(res.data.userInfo);
      }
    },

    async updateInfo(data: UserInfo) {
      const res = await updateUserInfo(data);
      this.setInfo(res.data.userInfo);
    },

    // Login
    async login(loginForm: LoginData) {
      try {
        console.log("LoginData:", loginForm);
        await userLogin(loginForm).then((res) => {
          console.log("res:", res);
          const { token, userInfo } = res;
          setToken(token);
          this.setInfo(userInfo);
        }).catch((e) => {
          console.log("e", e);
        });
      } catch (err) {
        clearToken();
        throw err;
      }
    },

    async loginMail(loginForm: LoginDataMail) {
      try {
        const res = await userLoginMail(loginForm);
        setToken(res.data.token);
      } catch (err) {
        clearToken();
        throw err;
      }
    },

    // Logout
    async logout() {
      this.resetInfo();
      clearToken();
      removeRouteListener();
    },
  },
});

export default useUserStore;