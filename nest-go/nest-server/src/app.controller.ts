import { Controller, Get } from '@nestjs/common';
import { EventPattern, MessagePattern } from '@nestjs/microservices';
import { AppService } from './app.service';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) { }
  @EventPattern('exchange')
  public async mock(data: any) {
    console.log('go-messsage', data);
  }

  @EventPattern('exchange-2')
  public async mock2(data: any) {
    console.log('go-messsage2', data);
  }

  @EventPattern('create-data-nats')
  public async mock3(data: any) {
    console.log('create-data-nats', data);
  }
  @Get()
  getHello(): string {
    return this.appService.getHello();
  }
}
