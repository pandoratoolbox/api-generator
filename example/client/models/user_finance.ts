import { IUser, User } from "./user";
import {
  IPaymentProviderSource,
  PaymentProviderSource,
} from "./payment_provider_source";

export interface IUserFinance {
  id?: number;
  user_id?: number;
  credits?: number;
  default_payment_provider_source_id?: number;
  user?: IUser;
  payment_provider_source?: IPaymentProviderSource;
}

export class UserFinance {
  id?: number;
  user_id?: number;
  credits?: number;
  default_payment_provider_source_id?: number;
  user?: User;
  payment_provider_source?: PaymentProviderSource;

  constructor(data: IUserFinance) {
    this.id = data.id;
    this.user_id = data.user_id;
    this.credits = data.credits;
    this.default_payment_provider_source_id =
      data.default_payment_provider_source_id;
    this.user = data.user ? new User(data.user) : undefined;
    this.payment_provider_source = data.payment_provider_source
      ? new PaymentProviderSource(data.payment_provider_source)
      : undefined;
  }
}
