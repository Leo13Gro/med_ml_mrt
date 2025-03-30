import ml_service.api.ml_api_pb2_grpc as pb_grpc
import ml_service.internal.usecases.mri.mri as mriusecase
from google.protobuf.empty_pb2 import Empty


class MlController(pb_grpc.MLAPIServicer):
    def __init__(self, mri_usecase: mriusecase.mriUseCase):
        super().__init__()
        self.mri_usecase = mri_usecase

    def SegmentAndClassification(self, request, context):
        print("запрос на сегментацию и классификацию")
        mri_id = request.mri_id

        try:
            self.mri_usecase.segmentAndClassificateByID(mri_id)
            return Empty()
        except Exception as e:
            context.set_details(f"Error processing request: {str(e)}")
            context.set_code(pb_grpc.StatusCode.INTERNAL)  # устанавливаем код ошибки
            return (
                None  # Возвращаем None в случае ошибки, чтобы завершить вызов с ошибкой
            )
