import {defineStore} from 'pinia';
import {Credentials} from 'components/models';
import AuthService from 'src/service/AuthService';
import { api } from 'boot/axios'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    token: '',
    username: '',
  }),
  persist: {
    afterRestore: (ctx) => {
      ctx.store.hydrate();
    }
  },
  getters: {
    isAuthenticated(): boolean {
      return this.token !== '';
    },
  },
  actions: {
    hydrate() {
      if (this.token) {
        api.defaults.headers.common['Authorization'] = 'Bearer ' + this.token
      }
    },
    logout() {
      this.token = '';
      this.username = '';
    },
    async login(credentials: Credentials): Promise<void> {
      return new Promise((resolve, reject) => {
        AuthService.login(credentials).then((res) => {
          api.defaults.headers.common['Authorization'] = 'Bearer ' + res.access_token
          this.token = res.access_token;
          resolve();
        }).catch((err) => reject(err));
      });
    },
    async register(credentials: Credentials): Promise<void> {
      return new Promise((resolve, reject) => {
        AuthService.register(credentials).then(() => {
          resolve();
        }).catch((err) => reject(err));
      });
    }
  },
});
