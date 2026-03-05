# aws-account-switcher

管理多个 AWS 账号本地 AK 的环境变量，支持一键切换和查询当前账号。

## 功能

- **管理多账号 AK**：本地保存多套 `AWS_ACCESS_KEY_ID` / `AWS_SECRET_ACCESS_KEY`（可选 `AWS_SESSION_TOKEN`），按账号名区分
- **一键切换**：执行一条命令后，当前 shell 的 AWS 环境变量切换到指定账号
- **查询当前账号**：根据当前环境变量显示正在使用的账号名

## 安装

```bash
cd aws-account-switcher
go build -o aws-account-switcher .
# 可选：放到 PATH
# mv aws-account-switcher ~/bin/ 或 sudo mv aws-account-switcher /usr/local/bin/
```

## 配置存储

- 路径：`~.aws-account-switcher/config.json`
- 权限：目录 0700，文件 0600（仅当前用户可读）

## 命令

| 命令            | 说明                                                   |
| --------------- | ------------------------------------------------------ |
| `add <name>`    | 添加或更新账号（交互输入 AK/SK，或从当前环境变量读取） |
| `list`          | 列出已保存的账号名（\* 表示当前配置的默认账号）        |
| `use <name>`    | 切换到指定账号（会输出供 `eval` 使用的命令）           |
| `export <name>` | 输出 shell 的 export 语句，供 `eval` 或 `source` 使用  |
| `current`       | 查询当前 shell 使用的账号                              |
| `remove <name>` | 删除已保存的账号                                       |

## 使用示例

### 添加账号

```bash
# 交互输入
./aws-account-switcher add prod

# 或先设置环境变量再添加（会优先使用环境变量）
export AWS_ACCESS_KEY_ID=AKIA...
export AWS_SECRET_ACCESS_KEY=...
./aws-account-switcher add prod
```

### 一键切换账号

```bash
# 方式一：eval 执行 export 输出
eval "$(./aws-account-switcher export prod)"

# 方式二：使用 use 会提示上述命令，并同时输出 export 到 stdout，也可 eval
eval "$(./aws-account-switcher use prod)"
```

### 查询当前账号

```bash
./aws-account-switcher current
# 输出示例：Current profile: prod
```

### 列出账号

```bash
./aws-account-switcher list
# 示例：
#   dev
# * prod
```

### 可选：设成别名

在 ~/.bashrc 或 ~/.zshrc 里加：

```sh
alias aws-use='function \_aws_use(){ eval "$(aws-account-switcher use $1)"; };_aws_use'
alias aws-use='function _aws_use(){ eval "$(aws-account-switcher use $1)"; };\_aws_use'
```

之后就可以用：

```sh
aws-use hix-awsaws-use pollo-aws
```

## 安全说明

- 密钥保存在本地 JSON 文件中，请勿提交到 Git
- 建议仅在本机使用，不要将 `accounts.json` 拷贝到其他环境
- 若需更高安全性，可考虑使用系统钥匙串或加密存储（需自行扩展）
