from ml_service.config.default import get_settings
from ml_service.internal.s3.s3 import S3
from ml_service.internal.ml_model.neuro_class import ModelABC
from ml_service.internal.ml_model.kt.neuro_class import KtModel
import ml_service.internal.events.kafka_pb2 as pb_event
import uuid

settings = get_settings()

class ktUseCase:
    def __init__(self, model: KtModel, store: S3):
        self.model = model
        self.store = store

    def segmentClassificateSave(self, kt_id):
        print("Going to S3...")
        data = self.store.load(kt_id + "/" + kt_id)

        result = self.model.predict(data)
        print(type(result))
        print(result)