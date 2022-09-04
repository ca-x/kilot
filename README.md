# kilot
tiny code generator  for quick development.
## functions
### current available
 + [x] generate mongodb code for kratos biz/data layer
```bash
NAME:
   kilot kratos mongo generate - start code generation.please pass at least one arg as model name(multi model names are supported).

USAGE:
   kilot kratos mongo generate [command options] [arguments...]

OPTIONS:
   --format, -f                                                                     set whether to format generated code before write to file. (default: true)
   --model-names ModelName, -n ModelName [ --model-names ModelName, -n ModelName ]  set model names ModelName for mongo code generate.multi model names supported.
   --model-output-dir Dir, -d Dir                                                   set model save dir Dir.if not set tool working dir will be used.
   --model-prefix Prefix, -p Prefix                                                 set model name prefix Prefix.default is empty.
   --model-suffix Suffix, -s Suffix                                                 set model name suffix Suffix .default is empty.
   --use-odm                                                                        this flag set whether to use mongo odm (use mgm). (default: true)

```
### ToDo
+ [ ] allow template custom
## install
run command to install
```bash
go install github.com/czyt/kilot/cmd/kilot@latest
```
## usage
after installed，run `kilot -h` for usage help.
