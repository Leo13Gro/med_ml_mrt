from confluent_kafka import Consumer
import ml_service.internal.events.kafka_pb2 as pb
from ml_service.config.default import get_settings

settings = get_settings()


class EventsYo:
    def __init__(self, mri):
        self.mri = mri

    def run(self):
        consumer_config = {
            "bootstrap.servers": settings.kafka_host + ":" + str(settings.kafka_port),  # Адрес Kafka-брокера
            "group.id": "mlService",  # Имя consumer group
            "auto.offset.reset": "earliest",  # Начинать с самого начала, если оффсет не найден
            "security.protocol": "PLAINTEXT",  # Установка протокола безопасности на PLAINTEXT для отключения SASL
            "broker.version.fallback": "2.3.0",
        }

        consumer = Consumer(consumer_config)
        consumer.subscribe(["mrisplitted"])
        while True:
            msg = consumer.poll(timeout=1.0)
            # continue
            if msg is None:
                continue  # Если сообщения нет, то пропускаем итерацию

            mri_splitted_event = pb.mriSplitted()
            mri_splitted_event.ParseFromString(msg.value())

            print("MRI ID: ", mri_splitted_event.mri_id)

            self.mri.segmentClassificateSave(
                mri_splitted_event.mri_id, mri_splitted_event.pages_id
            )
            consumer.commit(msg)
