import { useEffect, useState } from "react";
import api from "../../services/api";

function Team() {
  const [team, setTeam] = useState([]);
  const [members, setMembers] = useState([]);

  useEffect(() => {
    api.get("/read-team").then((r) => {
      setTeam(r.data);
    }).catch((e) => {
      if (e.response.status === 404) {
        setTeam({ title: "", content: "" });
      }
    });
  }, []);

  useEffect(() => {
    api.get("/read-members").then((r) => {
      setMembers(r.data.map((i) => {
        i.created = true;
        return i;
      }));
    }).catch((e) => {
      if (e.response.status === 404) {
        setMembers([]);
      }
    });
  }, []);

  const addMember = () => {
    setMembers(
      (
        prev,
      ) => ([...prev, {
        image: "",
        image_path: "",
        name: "",
        function: "",
        facebook: "",
        youtube: "",
        twitter: "",
        instagram: "",
      }]),
    );
  };

  const createMember = (i, index) => {
    if (i.created) {
      return;
    }
    const data = new FormData();

    data.append("image", i.image);
    data.append("name", i.name);
    data.append("function", i.function);
    data.append("facebook", i.facebook);
    data.append("youtube", i.youtube);
    data.append("twitter", i.twitter);
    data.append("instagram", i.instagram);

    api.post("/create-member", data, {
      headers: { "jwt-token": `JWT ${localStorage.getItem("token")}` },
    });

    setMembers((prev) =>
      prev.map((j, number) => {
        if (index === number) {
          j.created = true;
        }
        return j;
      })
    );
  };

  const deleteItem = (e, member) => {
    e.preventDefault();
    console.log(member);
    if (member.id === undefined) {
      setMembers(
        members.filter((i) =>
          JSON.stringify(i) === JSON.stringify(member) ? null : i
        ),
      );
      return;
    }
    api.delete("/delete-member/" + member.id, {
      headers: { "jwt-token": `JWT ${localStorage.getItem("token")}` },
    });

    setMembers(
      members.filter((i) =>
        JSON.stringify(i) === JSON.stringify(member) ? null : i
      ),
    );
  };

  const submitTeam = (e) => {
    e.preventDefault();
    api.post("/create-team", team, {
      headers: { "jwt-token": `JWT ${localStorage.getItem("token")}` },
    }).then();
  };

  if (Object.keys(team).length === 0 && members.length === 0) {
    return <h1>Loading ...</h1>;
  }

  return (<div
    style={{
      display: "flex",
      alignItems: "center",
      flexDirection: "column",
    }}
  >
    <h1>Team</h1>
    <form className={"Team"} onSubmit={submitTeam}>
      <div>
        <input
          type={"text"}
          placeholder={"Titulo"}
          value={team.title}
          onChange={(e) =>
            setTeam((prev) => ({ ...prev, title: e.target.value }))}
        />
      </div>
      <div>
        <input
          type={"text"}
          placeholder={"Conteudo"}
          value={team.content}
          onChange={(e) =>
            setTeam((prev) => ({ ...prev, content: e.target.value }))}
        />
      </div>
      <button type={"submit"}>
        Guardar alterações
      </button>
    </form>

    {members.map((i, index) => (
      <div key={index}>
        <hr />
        <h3>Membro</h3>
        <form>
          {Object.keys(i).map((item, n) => {
            if (item === "created" || item === "id") {
              return;
            }
            return (
              <div key={n}>
                {item === "image_path" && i[item] !== ""
                  ? <img
                    src={`${api.defaults.baseURL}${i[item]}`}
                    width="60px"
                    height="60px"
                  />
                  : null}
                {(() => {
                  if (item !== "image_path") {
                    return (<>
                      {item === "image"
                        ? <input
                          type={"file"}
                          value={i[item] === "" ? i[item] : i[item].filename}
                          onChange={(e) => {
                            setMembers(members.map((v, number) => {
                              if (index === number) {
                                v[item] = e.target.files[0];
                              }
                              return v;
                            }));
                          }}
                        />
                        : <input
                          placeholder={item}
                          value={i[item]}
                          onChange={((e) =>
                            setMembers((prev) =>
                              prev.map((member, number) => {
                                if (number === index) {
                                  member[item] = e.target.value;
                                }
                                return member;
                              })
                            ))}
                        />}
                    </>);
                  }
                })()}
              </div>
            );
          })}
          {i.created}
          {!i.created
            ? <div>
              <button type={"button"} onClick={() => createMember(i, index)}>
                Criar
              </button>
            </div>
            : null}
          <div>
            <button type={"button"} onClick={(e) => deleteItem(e, i)}>
              Apagar membro
            </button>
          </div>
        </form>
        <hr />
      </div>
    ))}
    <button type={"button"} onClick={addMember}>
      Adicionar membro
    </button>
  </div>);
}

export default Team;

/*


*/
