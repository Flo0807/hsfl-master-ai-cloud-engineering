import {defineStore} from 'pinia';
import BulletinBoardService from 'src/service/BulletinBoardService';
import {BulletinBoardEntry} from 'components/models';

export const useBulletinStore = defineStore('bulletin', {
  state: () => ({
    bulletins: [] as BulletinBoardEntry[],
    pagination: {
      currentPage: 0,
      pageSize: 0,
      totalPages: 0,
      totalRecords: 0,
    }
  }),
  getters: {
    get(): BulletinBoardEntry[] {
      return this.bulletins;
    },
    reachedEnd(): boolean {
      return this.pagination.currentPage >= this.pagination.totalPages;
    },
    isLoading(): boolean {
      return this.bulletins.length === 0;
    }
  },
  actions: {
    fetch() {
        const nextPage = this.pagination.currentPage + 1;
        BulletinBoardService.get(5, nextPage).then((res) => {
          this.bulletins.push(...res.records);
          this.pagination.currentPage = res.page.currentPage;
          this.pagination.pageSize = res.page.pageSize;
          this.pagination.totalPages = res.page.totalPages;
          this.pagination.totalRecords = res.page.totalRecords;
        });
    },
    create(bulletin: BulletinBoardEntry) {
      BulletinBoardService.create(bulletin).then((res) => {
        this.bulletins.unshift(res);
      });
    },
    update(bulletin: BulletinBoardEntry) {
      BulletinBoardService.update(bulletin.id, bulletin).then((res) => {
        const index = this.bulletins.findIndex((b) => b.id === res.id);
        this.bulletins[index] = res;
      });
    },
    remove(id: string) {
      BulletinBoardService.delete(id).then(() => {
        this.bulletins = this.bulletins.filter((b) => b.id !== id);
      });
    }
  }
});
