import { IPaymentProvider, PaymentProvider } from "./payment_provider";
import { IUser, User } from "./user";

export interface ISubscription {
  external_id?: number;
  user_id?: number;
  payment_provider?: IPaymentProvider;
  user?: IUser;
  id?: number;
  payment_provider_id?: number;
}

export class Subscription {
  external_id?: number;
  user_id?: number;
  payment_provider?: PaymentProvider;
  user?: User;
  id?: number;
  payment_provider_id?: number;

  constructor(data: ISubscription) {
    this.external_id = data.external_id;
    this.user_id = data.user_id;
    this.payment_provider = data.payment_provider
      ? new PaymentProvider(data.payment_provider)
      : undefined;
    this.user = data.user ? new User(data.user) : undefined;
    this.id = data.id;
    this.payment_provider_id = data.payment_provider_id;
  }
}
