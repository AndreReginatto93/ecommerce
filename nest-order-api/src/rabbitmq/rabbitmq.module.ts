/* eslint-disable @typescript-eslint/no-unsafe-call */
/* eslint-disable @typescript-eslint/no-unsafe-member-access */
import { RabbitMQModule } from '@golevelup/nestjs-rabbitmq';
import { Global, Module } from '@nestjs/common';

@Global()
@Module({
  imports: [
    RabbitMQModule.forRoot({
      uri: 'amqp://admin:admin@localhost:5672',
    }),
  ],
  exports: [RabbitMQModule],
})
export class RabbitmqModule {}
