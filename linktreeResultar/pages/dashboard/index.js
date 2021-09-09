import api from "../../services/api";
import AdminMenu from "../../components/AdminMenu/AdminMenu";

function Dashboard(props) {
  return (
    <div className={"dashboard"}>
      <AdminMenu file_path={props.file_path} />
      <div style={{ textAlign: "center", fontSize: "5rem" }}>Welcome</div>
    </div>
  );
}

Dashboard.getInitialProps = async () => {
  let file_path = undefined;
  file_path = await api.get("read-logo").then((r) => {
    return r.data.file_path;
  }).catch((e) => {
    if (e.response.status == 400) return undefined;
  });
  return { file_path: file_path };
};

export default Dashboard;
