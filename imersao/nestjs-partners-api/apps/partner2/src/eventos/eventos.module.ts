import { EventsCoreModule } from '@app/core/events/events-core.module';
import { Module } from '@nestjs/common';
import { EventosController } from './eventos.controller';

@Module({
  imports: [EventsCoreModule],
  controllers: [EventosController],
})
export class EventosModule {}
