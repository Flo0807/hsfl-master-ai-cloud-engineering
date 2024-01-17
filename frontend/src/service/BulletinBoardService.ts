import {BulletinBoardEntry, ResponsePage} from 'components/models';
import {api} from 'boot/axios';

const bulletinBoardApi = '/bulletin-board/posts';
export default {
  get(take: number, page: number): Promise<ResponsePage<BulletinBoardEntry>> {
    return new Promise((resolve, reject) => {
      api.get(`${bulletinBoardApi}`, {
        params: {
          take,
          page
        }
      }).then((res) => {
        if (res.status == 200) {
          resolve(res.data);
        } else {
          reject(new Error('Something went wrong'));
        }
      })
        .catch((err) => reject(err));
    });
  },
  create(bulletinBoardEntry: BulletinBoardEntry): Promise<BulletinBoardEntry> {
    return new Promise((resolve, reject) => {
      api.post(`${bulletinBoardApi}`, bulletinBoardEntry)
        .then((res) => {
          if (res.status == 201) {
            resolve(res.data);
          } else {
            reject(new Error('Something went wrong'));
          }
        })
        .catch((err) => reject(err));
    });
  },
  update(id: string, bulletinBoardEntry: BulletinBoardEntry): Promise<BulletinBoardEntry> {
    return new Promise((resolve, reject) => {
      api.put(`${bulletinBoardApi}/${id}`, bulletinBoardEntry)
        .then((res) => {
          if (res.status == 200) {
            resolve(res.data);
          } else {
            reject(new Error('Something went wrong'));
          }
        })
        .catch((err) => reject(err));
    });
  },
  delete(id: string): Promise<void> {
    return new Promise((resolve, reject) => {
      api.delete(`${bulletinBoardApi}/${id}`)
        .then((res) => {
          if (res.status == 200) {
            resolve();
          } else {
            reject(new Error('Something went wrong'));
          }
        })
        .catch((err) => reject(err));
    });
  }
}
