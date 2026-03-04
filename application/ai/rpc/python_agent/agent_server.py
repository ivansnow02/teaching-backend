"""
Python AI Agent — gRPC Server 参考实现
依赖: grpcio grpcio-tools langchain openai

生成 Python stub:
    python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. agent.proto
"""

import asyncio
from typing import AsyncIterator

import grpc
from grpc import aio

# 由 grpc_tools.protoc 从 agent.proto 生成
import agent_pb2
import agent_pb2_grpc

# ============================================================
# 示例: 使用 OpenAI / LangChain 进行流式推理
# ============================================================
# from langchain_openai import ChatOpenAI
# from langchain_core.messages import HumanMessage
# llm = ChatOpenAI(model="gpt-4o", streaming=True)


class AgentServicer(agent_pb2_grpc.AgentServiceServicer):
    """实现 agent.proto 中的 AgentService"""

    # ===================== 判卷 =====================

    async def GradeSubjective(self, request, context):
        """同步评分（Unary RPC）"""
        # 调用评分引擎（伪代码）
        # result = await grading_engine.grade(request)
        return agent_pb2.GradeSubjectiveRes(
            question_id=request.question_id,
            score="8",
            is_correct=2,       # 部分得分
            ai_comment="答案触及核心要点，但论证不够充分。",
        )

    # ===================== 课件向量化 =====================

    async def EmbedMaterial(self, request, context):
        """触发课件向量化（异步处理，立即返回 accepted）"""
        # await embedding_service.submit(request)
        return agent_pb2.EmbedMaterialRes(accepted=True)

    async def GetEmbedStatus(self, request, context):
        # status = await embedding_service.status(request.material_id)
        return agent_pb2.GetEmbedStatusRes(ai_status=2, message="已完成")

    # ===================== 智能助教 RAG =====================

    async def AskQuestion(self, request, context):
        """同步 RAG 问答（Unary RPC）"""
        return agent_pb2.AskQuestionRes(
            answer="微服务是一种将应用拆分为小型独立服务的架构模式。",
            sources=["chapter1.pdf#p3"],
        )

    async def AskQuestionStream(
        self, request, context
    ) -> AsyncIterator[agent_pb2.StreamDelta]:
        """
        流式 RAG 问答（Server-Side Streaming RPC）
        使用 `yield` 逐 token 推送，天然适配 gRPC Server Stream。
        """
        # === 实际场景使用 LangChain astream ===
        # async for chunk in llm.astream([HumanMessage(content=request.question)]):
        #     yield agent_pb2.StreamDelta(delta=chunk.content, finished=False)
        # yield agent_pb2.StreamDelta(delta="", sources=["..."], finished=True)

        # 演示：模拟逐字符流式输出
        answer = "微服务是一种将应用拆分为小型独立服务的架构模式，每个服务独立部署。"
        for char in answer:
            await asyncio.sleep(0.05)  # 模拟 LLM 推理延迟
            yield agent_pb2.StreamDelta(delta=char, finished=False)

        # 最后一帧携带引用来源
        yield agent_pb2.StreamDelta(
            delta="",
            sources=["chapter1.pdf#p3"],
            finished=True,
        )

    # ===================== 智能生成 =====================

    async def GenerateQuestions(self, request, context):
        # task_id = await task_queue.submit("generate_questions", request)
        return agent_pb2.GenerateTaskRes(task_id="task-abc-123")

    async def GenerateCourseOutline(self, request, context):
        return agent_pb2.GenerateTaskRes(task_id="task-outline-456")

    async def GenerateCourseware(self, request, context):
        return agent_pb2.GenerateTaskRes(task_id="task-cw-789")

    async def GenerateCoursewareStream(
        self, request, context
    ) -> AsyncIterator[agent_pb2.StreamDelta]:
        """
        流式生成课件 Markdown (Server-Side Streaming RPC)
        Python 侧 `yield` 每个 Markdown 片段。
        """
        # === 实际场景 ===
        # async for chunk in llm.astream(build_courseware_prompt(request)):
        #     yield agent_pb2.StreamDelta(delta=chunk.content, finished=False)
        # yield agent_pb2.StreamDelta(delta="", finished=True)

        markdown_pieces = [
            "# 微服务架构\n\n",
            "## 什么是微服务\n",
            "微服务是一种**架构风格**，将单一应用程序拆分为一组小型服务。\n\n",
            "## 核心优势\n",
            "- 独立部署\n- 技术异构\n- 弹性扩展\n",
        ]
        for piece in markdown_pieces:
            await asyncio.sleep(0.1)
            yield agent_pb2.StreamDelta(delta=piece, finished=False)

        yield agent_pb2.StreamDelta(delta="", finished=True)

    async def GetTaskStatus(self, request, context):
        # status = await task_queue.get_status(request.task_id)
        return agent_pb2.GetTaskStatusRes(
            status=1,  # 1: 已完成
            result='{"questions": [...]}',
            message="",
        )


async def serve():
    server = aio.server()
    agent_pb2_grpc.add_AgentServiceServicer_to_server(AgentServicer(), server)
    listen_addr = "[::]:50051"
    server.add_insecure_port(listen_addr)
    print(f"Python AI Agent gRPC Server started at {listen_addr}")
    await server.start()
    await server.wait_for_termination()


if __name__ == "__main__":
    asyncio.run(serve())
