import {
  IPaymentProviderCustomer,
  PaymentProviderCustomer,
} from "./payment_provider_customer";
import {
  IPaymentProviderSource,
  PaymentProviderSource,
} from "./payment_provider_source";
import { IInvoice, Invoice } from "./invoice";
import { ISubscription, Subscription } from "./subscription";

export interface IPaymentProvider {
  payment_provider_customer?: IPaymentProviderCustomer[];
  payment_provider_source?: IPaymentProviderSource[];
  id?: number;
  name?: string;
  is_enabled?: boolean;
  api_key?: string;
  invoice?: IInvoice[];
  subscription?: ISubscription[];
}

export class PaymentProvider {
  payment_provider_customer?: PaymentProviderCustomer[];
  payment_provider_source?: PaymentProviderSource[];
  id?: number;
  name?: string;
  is_enabled?: boolean;
  api_key?: string;
  invoice?: Invoice[];
  subscription?: Subscription[];

  constructor(data: IPaymentProvider) {
    this.payment_provider_customer = data.payment_provider_customer?.map(
      (i) => {
        return new PaymentProviderCustomer(i);
      }
    );
    this.payment_provider_source = data.payment_provider_source?.map((i) => {
      return new PaymentProviderSource(i);
    });
    this.id = data.id;
    this.name = data.name;
    this.is_enabled = data.is_enabled;
    this.api_key = data.api_key;
    this.invoice = data.invoice?.map((i) => {
      return new Invoice(i);
    });
    this.subscription = data.subscription?.map((i) => {
      return new Subscription(i);
    });
  }
}
