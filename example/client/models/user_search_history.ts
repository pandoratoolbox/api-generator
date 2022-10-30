import { IUser, User } from "./user";

export interface IUserSearchHistory {
  user_id?: number;
  user?: IUser;
  id?: number;
  last_search_at?: number | string | Date;
  queries?: string[];
}

export class UserSearchHistory {
  user_id?: number;
  user?: User;
  id?: number;
  last_search_at?: Date;
  queries?: string[];

  constructor(data: IUserSearchHistory) {
    this.user_id = data.user_id;
    this.user = data.user ? new User(data.user) : undefined;
    this.id = data.id;
    this.last_search_at = data.last_search_at
      ? new Date(data.last_search_at)
      : undefined;
    this.queries = data.queries;
  }
}
