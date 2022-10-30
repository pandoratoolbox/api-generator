import { IUser, User } from "./user";
import { IInvoice, Invoice } from "./invoice";
import { IDataPlatform, DataPlatform } from "./data_platform";

export interface IOrder {
  id?: number;
  scrape_job_id?: number;
  user?: IUser;
  created_at?: number | string | Date;
  invoice_id?: number;
  query?: string;
  invoice?: IInvoice;
  amount?: number;
  lead_ids?: number[];
  must_include_email?: boolean;
  user_id?: number;
  data_platform_id?: number;
  status?: string;
  data_platform?: IDataPlatform;
}

export class Order {
  id?: number;
  scrape_job_id?: number;
  user?: User;
  created_at?: Date;
  invoice_id?: number;
  query?: string;
  invoice?: Invoice;
  amount?: number;
  lead_ids?: number[];
  must_include_email?: boolean;
  user_id?: number;
  data_platform_id?: number;
  status?: string;
  data_platform?: DataPlatform;

  constructor(data: IOrder) {
    this.id = data.id;
    this.scrape_job_id = data.scrape_job_id;
    this.user = data.user ? new User(data.user) : undefined;
    this.created_at = data.created_at ? new Date(data.created_at) : undefined;
    this.invoice_id = data.invoice_id;
    this.query = data.query;
    this.invoice = data.invoice ? new Invoice(data.invoice) : undefined;
    this.amount = data.amount;
    this.lead_ids = data.lead_ids;
    this.must_include_email = data.must_include_email;
    this.user_id = data.user_id;
    this.data_platform_id = data.data_platform_id;
    this.status = data.status;
    this.data_platform = data.data_platform
      ? new DataPlatform(data.data_platform)
      : undefined;
  }
}
