import { IUser, User } from "./user";
import { IPaymentProvider, PaymentProvider } from "./payment_provider";
import { IUserFinance, UserFinance } from "./user_finance";

export interface IPaymentProviderSource {
  id?: number;
  user_id?: number;
  payment_provider_id?: number;
  external_id?: string;
  user?: IUser;
  payment_provider?: IPaymentProvider;
  user_finance?: IUserFinance[];
}

export class PaymentProviderSource {
  id?: number;
  user_id?: number;
  payment_provider_id?: number;
  external_id?: string;
  user?: User;
  payment_provider?: PaymentProvider;
  user_finance?: UserFinance[];

  constructor(data: IPaymentProviderSource) {
    this.id = data.id;
    this.user_id = data.user_id;
    this.payment_provider_id = data.payment_provider_id;
    this.external_id = data.external_id;
    this.user = data.user ? new User(data.user) : undefined;
    this.payment_provider = data.payment_provider
      ? new PaymentProvider(data.payment_provider)
      : undefined;
    this.user_finance = data.user_finance?.map((i) => {
      return new UserFinance(i);
    });
  }
}
