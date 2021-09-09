import { useEffect, useState } from "react";
import api from "../../services/api";

function Questions() {
  const [questionSection, setQuestionSection] = useState({});
  const [question, setQuestion] = useState([]);

  useEffect(() => {
    api.get("/read-questions").then((r) => {
      setQuestionSection(r.data);
    }).catch((e) => {
      if (e.response.status === 404) {
        setQuestionSection({ title: "", content: "" });
      }
    });
  }, []);

  useEffect(() => {
    api.get("/read-question").then((r) => {
      setQuestion(r.data.map((i) => {
        i.image = "";
        i.created = true;
        return i;
      }));
    }).catch((e) => {
      if (e.response.status === 404) {
        setQuestion([]);
      }
    });
  }, []);

  const handleSubmit = (e) => {
    e.preventDefault();
    api.post("/create-questions", questionSection, {
      headers: { "jwt-token": `JWT ${localStorage.getItem("token")}` },
    });
  };

  const createQuestion = (e, item) => {
    e.preventDefault();
    const data = new FormData();
    data.append("question", item.question);
    data.append("answer", item.answer);
    data.append("image", item.image);

    api.post("/create-question", data, {
      headers: { "jwt-token": `JWT ${localStorage.getItem("token")}` },
    });

    setQuestion(question.map((i) => {
      if (JSON.stringify(i) === JSON.stringify(item)) {
        i.created = true;
      }
      return i;
    }));
  };

  const deleteQuestion = (item) => {
    if (item.id === undefined) {
      setQuestion(
        question.filter((i) =>
          JSON.stringify(i) === JSON.stringify(item) ? null : i
        ),
      );
      return;
    }
    api.delete("/delete-question/" + item.id, {
      headers: { "jwt-token": `JWT ${localStorage.getItem("token")}` },
    });

    setQuestion(
      question.filter((i) =>
        JSON.stringify(i) === JSON.stringify(item) ? null : i
      ),
    );
  };

  const addQuestion = (e) => {
    setQuestion(
      (prev) => ([...prev, {
        question: "",
        answer: "",
        image: "",
        image_path: "",
      }]),
    );
  };

  if (Object.keys(questionSection).length === 0 && question.length === 0) {
    return <h1>Loading ...</h1>;
  }

  return (
    <div
      style={{
        display: "flex",
        alignItems: "center",
        flexDirection: "column",
      }}
    >
      <h1>Perguntas</h1>
      <form className={"questions"} onSubmit={handleSubmit}>
        <div>
          <input
            placeholder={"Titulo"}
            value={questionSection.title}
            onChange={(e) =>
              setQuestionSection((prev) => ({
                ...prev,
                title: e.target.value,
              }))}
          />
        </div>
        <div>
          <input
            placeholder={"Conteudo"}
            value={questionSection.content}
            onChange={(e) =>
              setQuestionSection((prev) => ({
                ...prev,
                content: e.target.value,
              }))}
          />
        </div>
        <button type={"submit"}>Guardar alterações</button>
      </form>
      {question.map((item, index) => {
        return (
          <div key={index}>
            <hr />
            <form>
              <div>
                <input
                  placeholder={"Pergunta"}
                  value={item.question}
                  onChange={(e) =>
                    setQuestion(
                      question.map((i, number) => {
                        if (number === index) {
                          i.question = e.target.value;
                        }
                        return i;
                      }),
                    )}
                />
              </div>
              <div>
                <input
                  placeholder={"Resposta"}
                  value={item.answer}
                  onChange={(e) =>
                    setQuestion(question.map((i, number) => {
                      if (number === index) {
                        i.answer = e.target.value;
                      }
                      return i;
                    }))}
                />
              </div>
              <div>
                {item.image_path === ""
                  ? (
                    <input
                      type={"file"}
                      value={item.image === ""
                        ? item.image
                        : item.image.filename}
                      onChange={(e) =>
                        setQuestion(
                          question.map((i, number) => {
                            if (number === index) {
                              i.image = e.target.files[0];
                            }
                            return i;
                          }),
                        )}
                    />
                  )
                  : (
                    <img
                      src={`${api.defaults.baseURL}${item.image_path}`}
                      width="60px"
                      height="60px"
                    />
                  )}
              </div>
              {!item.created
                ? (
                  <button
                    type={"button"}
                    onClick={(e) => createQuestion(e, item)}
                  >
                    Criar
                  </button>
                )
                : null}
              <button
                type={"button"}
                onClick={() => deleteQuestion(item, index)}
              >
                Apagar
              </button>
            </form>
            <hr />
          </div>
        );
      })}
      <button onClick={addQuestion}>Adicionar questão</button>
    </div>
  );
}

export default Questions;
