import { IPaymentProvider, PaymentProvider } from "./payment_provider";
import { IUser, User } from "./user";

export interface IPaymentProviderCustomer {
  payment_provider_id?: number;
  external_id?: string;
  user_id?: number;
  payment_provider?: IPaymentProvider;
  user?: IUser;
  id?: number;
}

export class PaymentProviderCustomer {
  payment_provider_id?: number;
  external_id?: string;
  user_id?: number;
  payment_provider?: PaymentProvider;
  user?: User;
  id?: number;

  constructor(data: IPaymentProviderCustomer) {
    this.payment_provider_id = data.payment_provider_id;
    this.external_id = data.external_id;
    this.user_id = data.user_id;
    this.payment_provider = data.payment_provider
      ? new PaymentProvider(data.payment_provider)
      : undefined;
    this.user = data.user ? new User(data.user) : undefined;
    this.id = data.id;
  }
}
