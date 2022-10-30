import { IUser, User } from "./user";
import { IInvoice, Invoice } from "./invoice";

export interface IOrderCredits {
  user?: IUser;
  id?: number;
  credits?: number;
  created_at?: number | string | Date;
  invoice_id?: number;
  user_id?: number;
  invoice?: IInvoice;
}

export class OrderCredits {
  user?: User;
  id?: number;
  credits?: number;
  created_at?: Date;
  invoice_id?: number;
  user_id?: number;
  invoice?: Invoice;

  constructor(data: IOrderCredits) {
    this.user = data.user ? new User(data.user) : undefined;
    this.id = data.id;
    this.credits = data.credits;
    this.created_at = data.created_at ? new Date(data.created_at) : undefined;
    this.invoice_id = data.invoice_id;
    this.user_id = data.user_id;
    this.invoice = data.invoice ? new Invoice(data.invoice) : undefined;
  }
}
