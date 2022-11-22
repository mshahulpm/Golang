import { NestFactory } from '@nestjs/core';
import { MicroserviceOptions, Transport } from '@nestjs/microservices';
import { AppModule } from './app.module';


async function bootstrap() {
  const app = await NestFactory.create(AppModule);

  app.connectMicroservice<MicroserviceOptions>({
    transport: Transport.RMQ,
    options: {
      urls: ['amqp://45.79.125.83:5672?heartbeat=30'],
      queue: 'exchange',
    },
  })

  await app.startAllMicroservices()

  await app.listen(4321);
}
bootstrap();
