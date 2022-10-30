import { IOrder, Order } from "./order";
import { IInvoice, Invoice } from "./invoice";
import { IOrderCredits, OrderCredits } from "./order_credits";
import { ISubscription, Subscription } from "./subscription";
import { IUserFinance, UserFinance } from "./user_finance";
import { IUser, User } from "./user";
import {
  IPaymentProviderCustomer,
  PaymentProviderCustomer,
} from "./payment_provider_customer";
import {
  IPaymentProviderSource,
  PaymentProviderSource,
} from "./payment_provider_source";
import { IUserSearchHistory, UserSearchHistory } from "./user_search_history";

export interface IUser {
  username?: string;
  order?: IOrder[];
  invoice?: IInvoice[];
  order_credits?: IOrderCredits[];
  id?: number;
  email?: string;
  referrer_id?: number;
  role_ids?: number[];
  subscription?: ISubscription[];
  user_finance?: IUserFinance;
  created_at?: number | string | Date;
  user?: IUser[];
  payment_provider_customer?: IPaymentProviderCustomer[];
  payment_provider_source?: IPaymentProviderSource[];
  password?: string;
  user_search_history?: IUserSearchHistory[];
}

export class User {
  username?: string;
  order?: Order[];
  invoice?: Invoice[];
  order_credits?: OrderCredits[];
  id?: number;
  email?: string;
  referrer_id?: number;
  role_ids?: number[];
  subscription?: Subscription[];
  user_finance?: UserFinance;
  created_at?: Date;
  user?: User[];
  payment_provider_customer?: PaymentProviderCustomer[];
  payment_provider_source?: PaymentProviderSource[];
  password?: string;
  user_search_history?: UserSearchHistory[];

  constructor(data: IUser) {
    this.username = data.username;
    this.order = data.order?.map((i) => {
      return new Order(i);
    });
    this.invoice = data.invoice?.map((i) => {
      return new Invoice(i);
    });
    this.order_credits = data.order_credits?.map((i) => {
      return new OrderCredits(i);
    });
    this.id = data.id;
    this.email = data.email;
    this.referrer_id = data.referrer_id;
    this.role_ids = data.role_ids;
    this.subscription = data.subscription?.map((i) => {
      return new Subscription(i);
    });
    this.user_finance = data.user_finance
      ? new UserFinance(data.user_finance)
      : undefined;
    this.created_at = data.created_at ? new Date(data.created_at) : undefined;
    this.user = data.user?.map((i) => {
      return new User(i);
    });
    this.payment_provider_customer = data.payment_provider_customer?.map(
      (i) => {
        return new PaymentProviderCustomer(i);
      }
    );
    this.payment_provider_source = data.payment_provider_source?.map((i) => {
      return new PaymentProviderSource(i);
    });
    this.password = data.password;
    this.user_search_history = data.user_search_history?.map((i) => {
      return new UserSearchHistory(i);
    });
  }
}
