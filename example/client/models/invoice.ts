import { IPaymentProvider, PaymentProvider } from "./payment_provider";
import { IOrder, Order } from "./order";
import { IOrderCredits, OrderCredits } from "./order_credits";
import { IUser, User } from "./user";

export interface IInvoice {
  id?: number;
  external_id?: string;
  user_id?: number;
  payment_provider?: IPaymentProvider;
  created_at?: number | string | Date;
  payment_provider_id?: number;
  amount_usd?: number;
  order?: IOrder;
  order_credits?: IOrderCredits;
  user?: IUser;
}

export class Invoice {
  id?: number;
  external_id?: string;
  user_id?: number;
  payment_provider?: PaymentProvider;
  created_at?: Date;
  payment_provider_id?: number;
  amount_usd?: number;
  order?: Order;
  order_credits?: OrderCredits;
  user?: User;

  constructor(data: IInvoice) {
    this.id = data.id;
    this.external_id = data.external_id;
    this.user_id = data.user_id;
    this.payment_provider = data.payment_provider
      ? new PaymentProvider(data.payment_provider)
      : undefined;
    this.created_at = data.created_at ? new Date(data.created_at) : undefined;
    this.payment_provider_id = data.payment_provider_id;
    this.amount_usd = data.amount_usd;
    this.order = data.order ? new Order(data.order) : undefined;
    this.order_credits = data.order_credits
      ? new OrderCredits(data.order_credits)
      : undefined;
    this.user = data.user ? new User(data.user) : undefined;
  }
}
