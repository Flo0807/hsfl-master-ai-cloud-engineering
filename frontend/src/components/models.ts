export interface Credentials {
  email: string;
  password: string;
}
export interface LoginResponse {
  access_token: string;
  token_type: string;
  expires_in: string;
}
export interface ResponsePage<T> {
  page: Pagination;
  records: T[];
}

export interface Pagination {
  currentPage: number;
  pageSize: number;
  totalPages: number;
  totalRecords: number;
}

export interface BulletinBoardEntry {
  id: string;
  title: string;
  content: string;
  createdAt: string;
}
