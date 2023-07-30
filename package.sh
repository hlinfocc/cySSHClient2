# compile for version
make
if [ $? -ne 0 ]; then
    echo "make error"
    exit 1
fi

cyssh_version=`./bin/cyssh -v`
echo "build version: $cyssh_version"

# cross_compiles
make -f ./Makefile.cross-compiles

rm -rf ./dist/packages
mkdir -p ./dist/packages

os_all='linux darwin freebsd'
arch_all='386 amd64 arm arm64 mips64 mips64le mips mipsle riscv64'

cd ./dist

for os in $os_all; do
    for arch in $arch_all; do
        cyssh_dir_name="cyssh-client_${cyssh_version}_${os}_${arch}"
        cyssh_path="./packages/cyssh-client_${cyssh_version}_${os}_${arch}"
        
        if [ ! -f "./cyssh_${os}_${arch}" ]; then
            continue
        fi
        if [ ! -f "./cyscp_${os}_${arch}" ]; then
            continue
        fi
        if [ ! -f "./cyssh-server_${os}_${arch}" ]; then
            continue
        fi
        mkdir ${cyssh_path}
        mv ./cyssh_${os}_${arch} ${cyssh_path}/cyssh
        mv ./cyscp_${os}_${arch} ${cyssh_path}/cyscp
        mv ./cyssh-server_${os}_${arch} ${cyssh_path}/cysh-server
        cp ../LICENSE ${cyssh_path}
        #cp -rf ../conf/* ${cyssh_path}

        # packages
        cd ./packages
        if [ "x${os}" = x"windows" ]; then
            zip -rq ${cyssh_dir_name}.zip ${cyssh_dir_name}
        else
            tar -zcvf ${cyssh_dir_name}.tar.gz ${cyssh_dir_name}
        fi  
        cd ..
        rm -rf ${cyssh_path}
    done
done

cd -
