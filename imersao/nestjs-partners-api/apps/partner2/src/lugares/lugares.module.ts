import { SpotsCoreModule } from '@app/core/spots/spots-core.module';
import { Module } from '@nestjs/common';
import { LugaresController } from './lugares.controller';

@Module({
  imports: [SpotsCoreModule],
  controllers: [LugaresController],
})
export class LugaresModule {}
