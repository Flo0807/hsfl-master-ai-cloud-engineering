import {api} from 'boot/axios';
import {Credentials, LoginResponse} from 'components/models';

const authApi = '/auth'

export default {
  login(credentials: Credentials): Promise<LoginResponse> {
    return new Promise((resolve, reject) => {
      api.post(`${authApi}/login`, credentials)
        .then((res) => {
          if (res.status == 200) {
            resolve(res.data);
          } else {
            reject(new Error('Something went wrong'));
          }
        }).catch((err) => reject(err));
    });
  },
  register(credentials: Credentials): Promise<void> {
    return new Promise((resolve, reject) => {
      api.post(`${authApi}/register`, credentials)
        .then((res) => {
          if (res.status == 201) {
            resolve();
          } else {
            reject(new Error('Something went wrong'));
          }
        }).catch((err) => reject(err));
    });
  }
}
