from confluent_kafka import Consumer
import ml_service.internal.events.kafka_pb2 as pb
from ml_service.config.default import get_settings

settings = get_settings()


class EventsYo:
    def __init__(self, uzi, kt):
        self.uzi = uzi
        self.kt = kt

    def run(self):
        consumer_config = {
            "bootstrap.servers": settings.kafka_host + ":" + str(settings.kafka_port),  # Адрес Kafka-брокера
            "group.id": "mlService",  # Имя consumer group
            "auto.offset.reset": "earliest",  # Начинать с самого начала, если оффсет не найден
            "security.protocol": "PLAINTEXT",  # Установка протокола безопасности на PLAINTEXT для отключения SASL
            "broker.version.fallback": "2.3.0",
        }

        consumer = Consumer(consumer_config)
        # Подписываемся на оба топика
        consumer.subscribe(["uzisplitted", "ktprepared"])
        while True:
            msg = consumer.poll(timeout=1.0)
            # continue
            if msg is None:
                continue  # Если сообщения нет, то пропускаем итерацию

            # Определяем обработчик в зависимости от топика
            print("topic:", msg.topic())
            if msg.topic() == "uzisplitted":
                self._process_uzi_message(msg)
            elif msg.topic() == "ktprepared":
                self._process_kt_message(msg)

            consumer.commit(msg)

    def _process_uzi_message(self, msg):
        uzi_splitted_event = pb.UziSplitted()
        print(msg.value())
        uzi_splitted_event.ParseFromString(msg.value())

        print("UZI ID: ", uzi_splitted_event.uzi_id)

        self.uzi.segmentClassificateSave(
            uzi_splitted_event.uzi_id, uzi_splitted_event.pages_id
        )

    def _process_kt_message(self, msg):
        kt_prepared_event = pb.KtPrepared()  # Предполагается, что у вас есть такой protobuf-класс
        kt_prepared_event.ParseFromString(msg.value())
        
        print("KT ID: ", kt_prepared_event.kt_id)

        self.kt.segmentClassificateSave(
            kt_prepared_event.kt_id
        )
        
        # Вызов соответствующего метода обработки
        # self.uzi.processKt(kt_prepared_event.kt_id, kt_prepared_event.data)