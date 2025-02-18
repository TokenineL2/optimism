// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const OptimismMintableERC721FactoryStorageLayoutJSON = "{\"storage\":[{\"astId\":27253,\"contract\":\"contracts/universal/OptimismMintableERC721Factory.sol:OptimismMintableERC721Factory\",\"label\":\"isOptimismMintableERC721\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_address,t_bool)\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_mapping(t_address,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_bool\"}}}"

var OptimismMintableERC721FactoryStorageLayout = new(solc.StorageLayout)

var OptimismMintableERC721FactoryDeployedBin = "0x60806040523480156200001157600080fd5b50600436106200006f5760003560e01c8063d97df6521162000056578063d97df65214620000cd578063e78cea92146200010a578063e9518196146200013257600080fd5b806354fd4d5014620000745780635572acae1462000096575b600080fd5b6200007e62000169565b6040516200008d9190620005dc565b60405180910390f35b620000bc620000a736600462000622565b60006020819052908152604090205460ff1681565b60405190151581526020016200008d565b620000e4620000de36600462000722565b62000214565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016200008d565b620000e47f000000000000000000000000000000000000000000000000000000000000000081565b6200015a7f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016200008d565b6060620001967f0000000000000000000000000000000000000000000000000000000000000000620003fa565b620001c17f0000000000000000000000000000000000000000000000000000000000000000620003fa565b620001ec7f0000000000000000000000000000000000000000000000000000000000000000620003fa565b60405160200162000200939291906200079f565b604051602081830303815290604052905090565b600073ffffffffffffffffffffffffffffffffffffffff8416620002e5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526044602482018190527f4f7074696d69736d4d696e7461626c65455243373231466163746f72793a204c908201527f3120746f6b656e20616464726573732063616e6e6f742062652061646472657360648201527f7328302900000000000000000000000000000000000000000000000000000000608482015260a40160405180910390fd5b60007f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000008686866040516200033a906200054f565b6200034a9594939291906200081b565b604051809103906000f08015801562000367573d6000803e3d6000fd5b5073ffffffffffffffffffffffffffffffffffffffff8181166000818152602081815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905590513381529394509188169290917fe72783bb8e0ca31286b85278da59684dd814df9762a52f0837f89edd1483b299910160405180910390a3949350505050565b6060816000036200043e57505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b81156200046e57806200045581620008ab565b9150620004669050600a8362000915565b915062000442565b60008167ffffffffffffffff8111156200048c576200048c62000640565b6040519080825280601f01601f191660200182016040528015620004b7576020820181803683370190505b5090505b84156200054757620004cf6001836200092c565b9150620004de600a8662000946565b620004eb9060306200095d565b60f81b81838151811062000503576200050362000978565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506200053f600a8662000915565b9450620004bb565b949350505050565b61311a80620009a883390190565b60005b838110156200057a57818101518382015260200162000560565b838111156200058a576000848401525b50505050565b60008151808452620005aa8160208601602086016200055d565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000620005f1602083018462000590565b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff811681146200061d57600080fd5b919050565b6000602082840312156200063557600080fd5b620005f182620005f8565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f8301126200068157600080fd5b813567ffffffffffffffff808211156200069f576200069f62000640565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715620006e857620006e862000640565b816040528381528660208588010111156200070257600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806000606084860312156200073857600080fd5b6200074384620005f8565b9250602084013567ffffffffffffffff808211156200076157600080fd5b6200076f878388016200066f565b935060408601359150808211156200078657600080fd5b5062000795868287016200066f565b9150509250925092565b60008451620007b38184602089016200055d565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551620007f1816001850160208a016200055d565b600192019182015283516200080e8160028401602088016200055d565b0160020195945050505050565b600073ffffffffffffffffffffffffffffffffffffffff808816835286602084015280861660408401525060a060608301526200085c60a083018562000590565b828103608084015262000870818562000590565b98975050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203620008df57620008df6200087c565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082620009275762000927620008e6565b500490565b6000828210156200094157620009416200087c565b500390565b600082620009585762000958620008e6565b500690565b600082198211156200097357620009736200087c565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfe60e06040523480156200001157600080fd5b506040516200311a3803806200311a83398101604081905262000034916200062d565b8181600062000044838262000756565b50600162000053828262000756565b5050506001600160a01b038516620000d85760405162461bcd60e51b815260206004820152603360248201527f4f7074696d69736d4d696e7461626c654552433732313a20627269646765206360448201527f616e6e6f7420626520616464726573732830290000000000000000000000000060648201526084015b60405180910390fd5b83600003620001505760405162461bcd60e51b815260206004820152603660248201527f4f7074696d69736d4d696e7461626c654552433732313a2072656d6f7465206360448201527f6861696e2069642063616e6e6f74206265207a65726f000000000000000000006064820152608401620000cf565b6001600160a01b038316620001ce5760405162461bcd60e51b815260206004820152603960248201527f4f7074696d69736d4d696e7461626c654552433732313a2072656d6f7465207460448201527f6f6b656e2063616e6e6f742062652061646472657373283029000000000000006064820152608401620000cf565b60808490526001600160a01b0383811660a081905290861660c0526200020290601462000256602090811b62000e6517901c565b62000218856200041660201b620010a81760201c565b6040516020016200022b92919062000822565b604051602081830303815290604052600a90816200024a919062000756565b50505050505062000993565b6060600062000267836002620008ac565b62000274906002620008ce565b6001600160401b038111156200028e576200028e62000553565b6040519080825280601f01601f191660200182016040528015620002b9576020820181803683370190505b509050600360fc1b81600081518110620002d757620002d7620008e9565b60200101906001600160f81b031916908160001a905350600f60fb1b81600181518110620003095762000309620008e9565b60200101906001600160f81b031916908160001a90535060006200032f846002620008ac565b6200033c906001620008ce565b90505b6001811115620003be576f181899199a1a9b1b9c1cb0b131b232b360811b85600f1660108110620003745762000374620008e9565b1a60f81b8282815181106200038d576200038d620008e9565b60200101906001600160f81b031916908160001a90535060049490941c93620003b681620008ff565b90506200033f565b5083156200040f5760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152606401620000cf565b9392505050565b6060816000036200043e5750506040805180820190915260018152600360fc1b602082015290565b8160005b81156200046e5780620004558162000919565b9150620004669050600a836200094b565b915062000442565b6000816001600160401b038111156200048b576200048b62000553565b6040519080825280601f01601f191660200182016040528015620004b6576020820181803683370190505b5090505b84156200052e57620004ce60018362000962565b9150620004dd600a866200097c565b620004ea906030620008ce565b60f81b818381518110620005025762000502620008e9565b60200101906001600160f81b031916908160001a90535062000526600a866200094b565b9450620004ba565b949350505050565b80516001600160a01b03811681146200054e57600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b60005b83811015620005865781810151838201526020016200056c565b8381111562000596576000848401525b50505050565b600082601f830112620005ae57600080fd5b81516001600160401b0380821115620005cb57620005cb62000553565b604051601f8301601f19908116603f01168101908282118183101715620005f657620005f662000553565b816040528381528660208588010111156200061057600080fd5b6200062384602083016020890162000569565b9695505050505050565b600080600080600060a086880312156200064657600080fd5b620006518662000536565b945060208601519350620006686040870162000536565b60608701519093506001600160401b03808211156200068657600080fd5b6200069489838a016200059c565b93506080880151915080821115620006ab57600080fd5b50620006ba888289016200059c565b9150509295509295909350565b600181811c90821680620006dc57607f821691505b602082108103620006fd57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200075157600081815260208120601f850160051c810160208610156200072c5750805b601f850160051c820191505b818110156200074d5782815560010162000738565b5050505b505050565b81516001600160401b0381111562000772576200077262000553565b6200078a81620007838454620006c7565b8462000703565b602080601f831160018114620007c25760008415620007a95750858301515b600019600386901b1c1916600185901b1785556200074d565b600085815260208120601f198616915b82811015620007f357888601518255948401946001909101908401620007d2565b5085821015620008125787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6832ba3432b932bab69d60b91b8152600083516200084881600985016020880162000569565b600160fe1b60099184019182015283516200086b81600a84016020880162000569565b712f746f6b656e5552493f75696e743235363d60701b600a9290910191820152601c01949350505050565b634e487b7160e01b600052601160045260246000fd5b6000816000190483118215151615620008c957620008c962000896565b500290565b60008219821115620008e457620008e462000896565b500190565b634e487b7160e01b600052603260045260246000fd5b60008162000911576200091162000896565b506000190190565b6000600182016200092e576200092e62000896565b5060010190565b634e487b7160e01b600052601260045260246000fd5b6000826200095d576200095d62000935565b500490565b60008282101562000977576200097762000896565b500390565b6000826200098e576200098e62000935565b500690565b60805160a05160c051612749620009d16000396000818161033001528181610a980152610bba015260006103090152600061035701526127496000f3fe608060405234801561001057600080fd5b50600436106101825760003560e01c806395d89b41116100d8578063c87b56dd1161008c578063e78cea9211610066578063e78cea921461032b578063e951819614610352578063e985e9c51461037957600080fd5b8063c87b56dd146102e9578063d547cfb7146102fc578063d6c0b2c41461030457600080fd5b8063a1448194116100bd578063a1448194146102b0578063a22cb465146102c3578063b88d4fde146102d657600080fd5b806395d89b41146102955780639dc29fac1461029d57600080fd5b806323b872dd1161013a5780634f6ccce7116101145780634f6ccce71461025c5780636352211e1461026f57806370a082311461028257600080fd5b806323b872dd146102235780632f745c591461023657806342842e0e1461024957600080fd5b8063081812fc1161016b578063081812fc146101c4578063095ea7b3146101fc57806318160ddd1461021157600080fd5b806301ffc9a71461018757806306fdde03146101af575b600080fd5b61019a610195366004612196565b6103c2565b60405190151581526020015b60405180910390f35b6101b7610471565b6040516101a69190612229565b6101d76101d236600461223c565b610503565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101a6565b61020f61020a36600461227e565b610537565b005b6008545b6040519081526020016101a6565b61020f6102313660046122a8565b6106c8565b61021561024436600461227e565b610769565b61020f6102573660046122a8565b610838565b61021561026a36600461223c565b610853565b6101d761027d36600461223c565b610911565b6102156102903660046122e4565b6109a3565b6101b7610a71565b61020f6102ab36600461227e565b610a80565b61020f6102be36600461227e565b610ba2565b61020f6102d13660046122ff565b610cb9565b61020f6102e436600461236a565b610cc8565b6101b76102f736600461223c565b610d70565b6101b7610dd7565b6101d77f000000000000000000000000000000000000000000000000000000000000000081565b6101d77f000000000000000000000000000000000000000000000000000000000000000081565b6102157f000000000000000000000000000000000000000000000000000000000000000081565b61019a610387366004612464565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260056020908152604080832093909416825291909152205460ff1690565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007fe49bc7f8000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000841682148061045a57507fffffffff00000000000000000000000000000000000000000000000000000000848116908216145b806104695750610469846111dd565b949350505050565b60606000805461048090612497565b80601f01602080910402602001604051908101604052809291908181526020018280546104ac90612497565b80156104f95780601f106104ce576101008083540402835291602001916104f9565b820191906000526020600020905b8154815290600101906020018083116104dc57829003601f168201915b5050505050905090565b600061050e82611233565b5060009081526004602052604090205473ffffffffffffffffffffffffffffffffffffffff1690565b600061054282610911565b90508073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610604576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e6560448201527f720000000000000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff8216148061062d575061062d8133610387565b6106b9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603e60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206e6f7220617070726f76656420666f7220616c6c000060648201526084016105fb565b6106c383836112c1565b505050565b6106d23382611361565b61075e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560448201527f72206e6f7220617070726f76656400000000000000000000000000000000000060648201526084016105fb565b6106c3838383611420565b6000610774836109a3565b8210610802576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f455243373231456e756d657261626c653a206f776e657220696e646578206f7560448201527f74206f6620626f756e647300000000000000000000000000000000000000000060648201526084016105fb565b5073ffffffffffffffffffffffffffffffffffffffff919091166000908152600660209081526040808320938352929052205490565b6106c383838360405180602001604052806000815250610cc8565b600061085e60085490565b82106108ec576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f455243373231456e756d657261626c653a20676c6f62616c20696e646578206f60448201527f7574206f6620626f756e6473000000000000000000000000000000000000000060648201526084016105fb565b600882815481106108ff576108ff6124ea565b90600052602060002001549050919050565b60008181526002602052604081205473ffffffffffffffffffffffffffffffffffffffff168061099d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f4552433732313a20696e76616c696420746f6b656e204944000000000000000060448201526064016105fb565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff8216610a48576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f74206120766160448201527f6c6964206f776e6572000000000000000000000000000000000000000000000060648201526084016105fb565b5073ffffffffffffffffffffffffffffffffffffffff1660009081526003602052604090205490565b60606001805461048090612497565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610b45576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f4f7074696d69736d4d696e7461626c654552433732313a206f6e6c792062726960448201527f6467652063616e2063616c6c20746869732066756e6374696f6e00000000000060648201526084016105fb565b610b4e81611692565b8173ffffffffffffffffffffffffffffffffffffffff167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca582604051610b9691815260200190565b60405180910390a25050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610c67576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f4f7074696d69736d4d696e7461626c654552433732313a206f6e6c792062726960448201527f6467652063616e2063616c6c20746869732066756e6374696f6e00000000000060648201526084016105fb565b610c71828261176b565b8173ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d412139688582604051610b9691815260200190565b610cc4338383611785565b5050565b610cd23383611361565b610d5e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560448201527f72206e6f7220617070726f76656400000000000000000000000000000000000060648201526084016105fb565b610d6a848484846118b2565b50505050565b6060610d7b82611233565b6000610d85611955565b90506000815111610da55760405180602001604052806000815250610dd0565b80610daf846110a8565b604051602001610dc0929190612519565b6040516020818303038152906040525b9392505050565b600a8054610de490612497565b80601f0160208091040260200160405190810160405280929190818152602001828054610e1090612497565b8015610e5d5780601f10610e3257610100808354040283529160200191610e5d565b820191906000526020600020905b815481529060010190602001808311610e4057829003601f168201915b505050505081565b60606000610e74836002612577565b610e7f9060026125b4565b67ffffffffffffffff811115610e9757610e9761233b565b6040519080825280601f01601f191660200182016040528015610ec1576020820181803683370190505b5090507f300000000000000000000000000000000000000000000000000000000000000081600081518110610ef857610ef86124ea565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110610f5b57610f5b6124ea565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506000610f97846002612577565b610fa29060016125b4565b90505b600181111561103f577f303132333435363738396162636465660000000000000000000000000000000085600f1660108110610fe357610fe36124ea565b1a60f81b828281518110610ff957610ff96124ea565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060049490941c93611038816125cc565b9050610fa5565b508315610dd0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e7460448201526064016105fb565b6060816000036110eb57505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b811561111557806110ff81612601565b915061110e9050600a83612668565b91506110ef565b60008167ffffffffffffffff8111156111305761113061233b565b6040519080825280601f01601f19166020018201604052801561115a576020820181803683370190505b5090505b84156104695761116f60018361267c565b915061117c600a86612693565b6111879060306125b4565b60f81b81838151811061119c5761119c6124ea565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506111d6600a86612668565b945061115e565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f780e9d6300000000000000000000000000000000000000000000000000000000148061099d575061099d82611964565b60008181526002602052604090205473ffffffffffffffffffffffffffffffffffffffff166112be576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f4552433732313a20696e76616c696420746f6b656e204944000000000000000060448201526064016105fb565b50565b600081815260046020526040902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8416908117909155819061131b82610911565b73ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b60008061136d83610911565b90508073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614806113db575073ffffffffffffffffffffffffffffffffffffffff80821660009081526005602090815260408083209388168352929052205460ff165b8061046957508373ffffffffffffffffffffffffffffffffffffffff1661140184610503565b73ffffffffffffffffffffffffffffffffffffffff1614949350505050565b8273ffffffffffffffffffffffffffffffffffffffff1661144082610911565b73ffffffffffffffffffffffffffffffffffffffff16146114e3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060448201527f6f776e657200000000000000000000000000000000000000000000000000000060648201526084016105fb565b73ffffffffffffffffffffffffffffffffffffffff8216611585576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f2061646460448201527f726573730000000000000000000000000000000000000000000000000000000060648201526084016105fb565b611590838383611a47565b61159b6000826112c1565b73ffffffffffffffffffffffffffffffffffffffff831660009081526003602052604081208054600192906115d190849061267c565b909155505073ffffffffffffffffffffffffffffffffffffffff8216600090815260036020526040812080546001929061160c9084906125b4565b909155505060008181526002602052604080822080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff86811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b600061169d82610911565b90506116ab81600084611a47565b6116b66000836112c1565b73ffffffffffffffffffffffffffffffffffffffff811660009081526003602052604081208054600192906116ec90849061267c565b909155505060008281526002602052604080822080547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555183919073ffffffffffffffffffffffffffffffffffffffff8416907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908390a45050565b610cc4828260405180602001604052806000815250611b4d565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361181a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c65720000000000000060448201526064016105fb565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526005602090815260408083209487168084529482529182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6118bd848484611420565b6118c984848484611bf0565b610d6a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603260248201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560448201527f63656976657220696d706c656d656e746572000000000000000000000000000060648201526084016105fb565b6060600a805461048090612497565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f80ac58cd0000000000000000000000000000000000000000000000000000000014806119f757507fffffffff0000000000000000000000000000000000000000000000000000000082167f5b5e139f00000000000000000000000000000000000000000000000000000000145b8061099d57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000083161461099d565b73ffffffffffffffffffffffffffffffffffffffff8316611aaf57611aaa81600880546000838152600960205260408120829055600182018355919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30155565b611aec565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614611aec57611aec8382611de3565b73ffffffffffffffffffffffffffffffffffffffff8216611b10576106c381611e9a565b8273ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16146106c3576106c38282611f49565b611b578383611f9a565b611b646000848484611bf0565b6106c3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603260248201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560448201527f63656976657220696d706c656d656e746572000000000000000000000000000060648201526084016105fb565b600073ffffffffffffffffffffffffffffffffffffffff84163b15611dd8576040517f150b7a0200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85169063150b7a0290611c679033908990889088906004016126a7565b6020604051808303816000875af1925050508015611cc0575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201909252611cbd918101906126f0565b60015b611d8d573d808015611cee576040519150601f19603f3d011682016040523d82523d6000602084013e611cf3565b606091505b508051600003611d85576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603260248201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560448201527f63656976657220696d706c656d656e746572000000000000000000000000000060648201526084016105fb565b805181602001fd5b7fffffffff00000000000000000000000000000000000000000000000000000000167f150b7a0200000000000000000000000000000000000000000000000000000000149050610469565b506001949350505050565b60006001611df0846109a3565b611dfa919061267c565b600083815260076020526040902054909150808214611e5a5773ffffffffffffffffffffffffffffffffffffffff841660009081526006602090815260408083208584528252808320548484528184208190558352600790915290208190555b50600091825260076020908152604080842084905573ffffffffffffffffffffffffffffffffffffffff9094168352600681528383209183525290812055565b600854600090611eac9060019061267c565b60008381526009602052604081205460088054939450909284908110611ed457611ed46124ea565b906000526020600020015490508060088381548110611ef557611ef56124ea565b6000918252602080832090910192909255828152600990915260408082208490558582528120556008805480611f2d57611f2d61270d565b6001900381819060005260206000200160009055905550505050565b6000611f54836109a3565b73ffffffffffffffffffffffffffffffffffffffff9093166000908152600660209081526040808320868452825280832085905593825260079052919091209190915550565b73ffffffffffffffffffffffffffffffffffffffff8216612017576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f206164647265737360448201526064016105fb565b60008181526002602052604090205473ffffffffffffffffffffffffffffffffffffffff16156120a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e7465640000000060448201526064016105fb565b6120af60008383611a47565b73ffffffffffffffffffffffffffffffffffffffff821660009081526003602052604081208054600192906120e59084906125b4565b909155505060008181526002602052604080822080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff861690811790915590518392907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b7fffffffff00000000000000000000000000000000000000000000000000000000811681146112be57600080fd5b6000602082840312156121a857600080fd5b8135610dd081612168565b60005b838110156121ce5781810151838201526020016121b6565b83811115610d6a5750506000910152565b600081518084526121f78160208601602086016121b3565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610dd060208301846121df565b60006020828403121561224e57600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461227957600080fd5b919050565b6000806040838503121561229157600080fd5b61229a83612255565b946020939093013593505050565b6000806000606084860312156122bd57600080fd5b6122c684612255565b92506122d460208501612255565b9150604084013590509250925092565b6000602082840312156122f657600080fd5b610dd082612255565b6000806040838503121561231257600080fd5b61231b83612255565b91506020830135801515811461233057600080fd5b809150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000806000806080858703121561238057600080fd5b61238985612255565b935061239760208601612255565b925060408501359150606085013567ffffffffffffffff808211156123bb57600080fd5b818701915087601f8301126123cf57600080fd5b8135818111156123e1576123e161233b565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019083821181831017156124275761242761233b565b816040528281528a602084870101111561244057600080fd5b82602086016020830137600060208483010152809550505050505092959194509250565b6000806040838503121561247757600080fd5b61248083612255565b915061248e60208401612255565b90509250929050565b600181811c908216806124ab57607f821691505b6020821081036124e4577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000835161252b8184602088016121b3565b83519083019061253f8183602088016121b3565b01949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156125af576125af612548565b500290565b600082198211156125c7576125c7612548565b500190565b6000816125db576125db612548565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361263257612632612548565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60008261267757612677612639565b500490565b60008282101561268e5761268e612548565b500390565b6000826126a2576126a2612639565b500690565b600073ffffffffffffffffffffffffffffffffffffffff8087168352808616602084015250836040830152608060608301526126e660808301846121df565b9695505050505050565b60006020828403121561270257600080fd5b8151610dd081612168565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea164736f6c634300080f000aa164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(OptimismMintableERC721FactoryStorageLayoutJSON), OptimismMintableERC721FactoryStorageLayout); err != nil {
		panic(err)
	}

	layouts["OptimismMintableERC721Factory"] = OptimismMintableERC721FactoryStorageLayout
	deployedBytecodes["OptimismMintableERC721Factory"] = OptimismMintableERC721FactoryDeployedBin
}
